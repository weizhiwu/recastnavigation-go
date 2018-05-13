package detour

import (
	"math"
)

/**
@defgroup detour Detour

Members in this module are used to create, manipulate, and query navigation
meshes.

@note This is a summary list of members.  Use the index or search
feature to find minor members.
*/

/// @name General helper functions
/// @{

/// Used to ignore a function parameter.  VS complains about unused parameters
/// and this silences the warning.
///  @param [in] _ Unused parameter
func DtIgnoreUnused(interface{}) {}

/// Swaps the values of the two parameters.
///  @param[in,out]	a	Value A
///  @param[in,out]	b	Value B
func DtSwapFloat64(a, b *float64) { t := *a; *a = *b; *b = t }
func DtSwapUInt32(a, b *uint32)   { t := *a; *a = *b; *b = t }
func DtSwapInt32(a, b *int32)     { t := *a; *a = *b; *b = t }
func DtSwapUInt16(a, b *uint16)   { t := *a; *a = *b; *b = t }
func DtSwapInt16(a, b *int16)     { t := *a; *a = *b; *b = t }

/// Returns the minimum of two values.
///  @param[in]		a	Value A
///  @param[in]		b	Value B
///  @return The minimum of the two values.
func DtMinFloat64(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}
func DtMinUInt32(a, b uint32) uint32 {
	if a < b {
		return a
	} else {
		return b
	}
}
func DtMinInt32(a, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}
func DtMinUInt16(a, b uint16) uint16 {
	if a < b {
		return a
	} else {
		return b
	}
}
func DtMinInt16(a, b int16) int16 {
	if a < b {
		return a
	} else {
		return b
	}
}

/// Returns the maximum of two values.
///  @param[in]		a	Value A
///  @param[in]		b	Value B
///  @return The maximum of the two values.
func DtMaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}
func DtMaxUInt32(a, b uint32) uint32 {
	if a > b {
		return a
	} else {
		return b
	}
}
func DtMaxInt32(a, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}
func DtMaxUInt16(a, b uint16) uint16 {
	if a > b {
		return a
	} else {
		return b
	}
}
func DtMaxInt16(a, b int16) int16 {
	if a > b {
		return a
	} else {
		return b
	}
}

/// Returns the absolute value.
///  @param[in]		a	The value.
///  @return The absolute value of the specified value.
func DtAbsFloat64(a float64) float64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
func DtAbsInt32(a int32) int32 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
func DtAbsInt16(a int16) int16 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

/// Returns the square of the value.
///  @param[in]		a	The value.
///  @return The square of the value.
func DtSqrFloat64(a float64) float64 { return a * a }
func DtSqrUInt32(a uint32) uint32    { return a * a }
func DtSqrInt32(a int32) int32       { return a * a }
func DtSqrUInt16(a uint16) uint16    { return a * a }
func DtSqrInt16(a int16) int16       { return a * a }

/// Clamps the value to the specified range.
///  @param[in]		v	The value to clamp.
///  @param[in]		mn	The minimum permitted return value.
///  @param[in]		mx	The maximum permitted return value.
///  @return The value, clamped to the specified range.
func DtClampFloat64(v, mn, mx float64) float64 {
	if v < mn {
		return mn
	} else {
		if v > mx {
			return mx
		} else {
			return v
		}
	}
}
func DtClampUInt32(v, mn, mx uint32) uint32 {
	if v < mn {
		return mn
	} else {
		if v > mx {
			return mx
		} else {
			return v
		}
	}
}
func DtClampInt32(v, mn, mx int32) int32 {
	if v < mn {
		return mn
	} else {
		if v > mx {
			return mx
		} else {
			return v
		}
	}
}
func DtClampUInt16(v, mn, mx uint16) uint16 {
	if v < mn {
		return mn
	} else {
		if v > mx {
			return mx
		} else {
			return v
		}
	}
}
func DtClampInt16(v, mn, mx int16) int16 {
	if v < mn {
		return mn
	} else {
		if v > mx {
			return mx
		} else {
			return v
		}
	}
}

/// @}
/// @name Vector helper functions.
/// @{

/// Derives the cross product of two vectors. (@p v1 x @p v2)
///  @param[out]	dest	The cross product. [(x, y, z)]
///  @param[in]		v1		A Vector [(x, y, z)]
///  @param[in]		v2		A vector [(x, y, z)]
func DtVcross(dest, v1, v2 []float64) {
	dest[0] = v1[1]*v2[2] - v1[2]*v2[1]
	dest[1] = v1[2]*v2[0] - v1[0]*v2[2]
	dest[2] = v1[0]*v2[1] - v1[1]*v2[0]
}

/// Derives the dot product of two vectors. (@p v1 . @p v2)
///  @param[in]		v1	A Vector [(x, y, z)]
///  @param[in]		v2	A vector [(x, y, z)]
/// @return The dot product.
func DtVdot(v1, v2 []float64) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

/// Performs a scaled vector addition. (@p v1 + (@p v2 * @p s))
///  @param[out]	dest	The result vector. [(x, y, z)]
///  @param[in]		v1		The base vector. [(x, y, z)]
///  @param[in]		v2		The vector to scale and add to @p v1. [(x, y, z)]
///  @param[in]		s		The amount to scale @p v2 by before adding to @p v1.
func DtVmad(dest, v1, v2 []float64, s float64) {
	dest[0] = v1[0] + v2[0]*s
	dest[1] = v1[1] + v2[1]*s
	dest[2] = v1[2] + v2[2]*s
}

/// Performs a linear interpolation between two vectors. (@p v1 toward @p v2)
///  @param[out]	dest	The result vector. [(x, y, x)]
///  @param[in]		v1		The starting vector.
///  @param[in]		v2		The destination vector.
///	 @param[in]		t		The interpolation factor. [Limits: 0 <= value <= 1.0]
func DtVlerp(dest, v1, v2 []float64, t float64) {
	dest[0] = v1[0] + (v2[0]-v1[0])*t
	dest[1] = v1[1] + (v2[1]-v1[1])*t
	dest[2] = v1[2] + (v2[2]-v1[2])*t
}

/// Performs a vector addition. (@p v1 + @p v2)
///  @param[out]	dest	The result vector. [(x, y, z)]
///  @param[in]		v1		The base vector. [(x, y, z)]
///  @param[in]		v2		The vector to add to @p v1. [(x, y, z)]
func DtVadd(dest, v1, v2 []float64) {
	dest[0] = v1[0] + v2[0]
	dest[1] = v1[1] + v2[1]
	dest[2] = v1[2] + v2[2]
}

/// Performs a vector subtraction. (@p v1 - @p v2)
///  @param[out]	dest	The result vector. [(x, y, z)]
///  @param[in]		v1		The base vector. [(x, y, z)]
///  @param[in]		v2		The vector to subtract from @p v1. [(x, y, z)]
func DtVsub(dest, v1, v2 []float64) {
	dest[0] = v1[0] - v2[0]
	dest[1] = v1[1] - v2[1]
	dest[2] = v1[2] - v2[2]
}

/// Scales the vector by the specified value. (@p v * @p t)
///  @param[out]	dest	The result vector. [(x, y, z)]
///  @param[in]		v		The vector to scale. [(x, y, z)]
///  @param[in]		t		The scaling factor.
func DtVscale(dest, v []float64, t float64) {
	dest[0] = v[0] * t
	dest[1] = v[1] * t
	dest[2] = v[2] * t
}

/// Selects the minimum value of each element from the specified vectors.
///  @param[in,out]	mn	A vector.  (Will be updated with the result.) [(x, y, z)]
///  @param[in]	v	A vector. [(x, y, z)]
func DtVmin(mn, v []float64) {
	mn[0] = DtMinFloat64(mn[0], v[0])
	mn[1] = DtMinFloat64(mn[1], v[1])
	mn[2] = DtMinFloat64(mn[2], v[2])
}

/// Selects the maximum value of each element from the specified vectors.
///  @param[in,out]	mx	A vector.  (Will be updated with the result.) [(x, y, z)]
///  @param[in]		v	A vector. [(x, y, z)]
func DtVmax(mx, v []float64) {
	mx[0] = DtMaxFloat64(mx[0], v[0])
	mx[1] = DtMaxFloat64(mx[1], v[1])
	mx[2] = DtMaxFloat64(mx[2], v[2])
}

/// Sets the vector elements to the specified values.
///  @param[out]	dest	The result vector. [(x, y, z)]
///  @param[in]		x		The x-value of the vector.
///  @param[in]		y		The y-value of the vector.
///  @param[in]		z		The z-value of the vector.
func DtVset(dest []float64, x, y, z float64) {
	dest[0] = x
	dest[1] = y
	dest[2] = z
}

/// Performs a vector copy.
///  @param[out]	dest	The result. [(x, y, z)]
///  @param[in]		a		The vector to copy. [(x, y, z)]
func DtVcopy(dest, a []float64) {
	dest[0] = a[0]
	dest[1] = a[1]
	dest[2] = a[2]
}

/// Derives the scalar length of the vector.
///  @param[in]		v The vector. [(x, y, z)]
/// @return The scalar length of the vector.
func DtVlen(v []float64) float64 {
	return DtMathSqrtf(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

/// Derives the square of the scalar length of the vector. (len * len)
///  @param[in]		v The vector. [(x, y, z)]
/// @return The square of the scalar length of the vector.
func DtVlenSqr(v []float64) float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

/// Returns the distance between two points.
///  @param[in]		v1	A point. [(x, y, z)]
///  @param[in]		v2	A point. [(x, y, z)]
/// @return The distance between the two points.
func DtVdist(v1, v2 []float64) float64 {
	dx := v2[0] - v1[0]
	dy := v2[1] - v1[1]
	dz := v2[2] - v1[2]
	return DtMathSqrtf(dx*dx + dy*dy + dz*dz)
}

/// Returns the square of the distance between two points.
///  @param[in]		v1	A point. [(x, y, z)]
///  @param[in]		v2	A point. [(x, y, z)]
/// @return The square of the distance between the two points.
func DtVdistSqr(v1, v2 []float64) float64 {
	dx := v2[0] - v1[0]
	dy := v2[1] - v1[1]
	dz := v2[2] - v1[2]
	return dx*dx + dy*dy + dz*dz
}

/// Derives the distance between the specified points on the xz-plane.
///  @param[in]		v1	A point. [(x, y, z)]
///  @param[in]		v2	A point. [(x, y, z)]
/// @return The distance between the point on the xz-plane.
///
/// The vectors are projected onto the xz-plane, so the y-values are ignored.
func DtVdist2D(v1, v2 []float64) float64 {
	dx := v2[0] - v1[0]
	dz := v2[2] - v1[2]
	return DtMathSqrtf(dx*dx + dz*dz)
}

/// Derives the square of the distance between the specified points on the xz-plane.
///  @param[in]		v1	A point. [(x, y, z)]
///  @param[in]		v2	A point. [(x, y, z)]
/// @return The square of the distance between the point on the xz-plane.
func DtVdist2DSqr(v1, v2 []float64) float64 {
	dx := v2[0] - v1[0]
	dz := v2[2] - v1[2]
	return dx*dx + dz*dz
}

/// Normalizes the vector.
///  @param[in,out]	v	The vector to normalize. [(x, y, z)]
func DtVnormalize(v []float64) {
	d := 1.0 / DtMathSqrtf(DtSqrFloat64(v[0])+DtSqrFloat64(v[1])+DtSqrFloat64(v[2]))
	v[0] *= d
	v[1] *= d
	v[2] *= d
}

var thr float64 = DtSqrFloat64(1.0 / 16384.0)

/// Performs a 'sloppy' colocation check of the specified points.
///  @param[in]		p0	A point. [(x, y, z)]
///  @param[in]		p1	A point. [(x, y, z)]
/// @return True if the points are considered to be at the same location.
///
/// Basically, this function will return true if the specified points are
/// close enough to eachother to be considered colocated.
func DtVequal(p0, p1 []float64) bool {
	d := DtVdistSqr(p0, p1)
	return d < thr
}

/// Derives the dot product of two vectors on the xz-plane. (@p u . @p v)
///  @param[in]		u		A vector [(x, y, z)]
///  @param[in]		v		A vector [(x, y, z)]
/// @return The dot product on the xz-plane.
///
/// The vectors are projected onto the xz-plane, so the y-values are ignored.
func DtVdot2D(u, v []float64) float64 {
	return u[0]*v[0] + u[2]*v[2]
}

/// Derives the xz-plane 2D perp product of the two vectors. (uz*vx - ux*vz)
///  @param[in]		u		The LHV vector [(x, y, z)]
///  @param[in]		v		The RHV vector [(x, y, z)]
/// @return The dot product on the xz-plane.
///
/// The vectors are projected onto the xz-plane, so the y-values are ignored.
func DtVperp2D(u, v []float64) float64 {
	return u[2]*v[0] - u[0]*v[2]
}

/// @}
/// @name Computational geometry helper functions.
/// @{

/// Derives the signed xz-plane area of the triangle ABC, or the relationship of line AB to point C.
///  @param[in]		a		Vertex A. [(x, y, z)]
///  @param[in]		b		Vertex B. [(x, y, z)]
///  @param[in]		c		Vertex C. [(x, y, z)]
/// @return The signed xz-plane area of the triangle.
func DtTriArea2D(a, b, c []float64) float64 {
	abx := b[0] - a[0]
	abz := b[2] - a[2]
	acx := c[0] - a[0]
	acz := c[2] - a[2]
	return acx*abz - abx*acz
}

/// Determines if two axis-aligned bounding boxes overlap.
///  @param[in]		amin	Minimum bounds of box A. [(x, y, z)]
///  @param[in]		amax	Maximum bounds of box A. [(x, y, z)]
///  @param[in]		bmin	Minimum bounds of box B. [(x, y, z)]
///  @param[in]		bmax	Maximum bounds of box B. [(x, y, z)]
/// @return True if the two AABB's overlap.
/// @see dtOverlapBounds
func DtOverlapQuantBounds(amin, amax, bmin, bmax []uint16) bool {
	return !(amin[0] > bmax[0] || amax[0] < bmin[0] || amin[1] > bmax[1] || amax[1] < bmin[1] || amin[2] > bmax[2] || amax[2] < bmin[2])
}

/// Determines if two axis-aligned bounding boxes overlap.
///  @param[in]		amin	Minimum bounds of box A. [(x, y, z)]
///  @param[in]		amax	Maximum bounds of box A. [(x, y, z)]
///  @param[in]		bmin	Minimum bounds of box B. [(x, y, z)]
///  @param[in]		bmax	Maximum bounds of box B. [(x, y, z)]
/// @return True if the two AABB's overlap.
/// @see dtOverlapQuantBounds
func DtOverlapBounds(amin, amax, bmin, bmax []float64) bool {
	return !(amin[0] > bmax[0] || amax[0] < bmin[0] || amin[1] > bmax[1] || amax[1] < bmin[1] || amin[2] > bmax[2] || amax[2] < bmin[2])
}

/// Derives the closest point on a triangle from the specified reference point.
///  @param[out]	closest	The closest point on the triangle.
///  @param[in]		p		The reference point from which to test. [(x, y, z)]
///  @param[in]		a		Vertex A of triangle ABC. [(x, y, z)]
///  @param[in]		b		Vertex B of triangle ABC. [(x, y, z)]
///  @param[in]		c		Vertex C of triangle ABC. [(x, y, z)]
func DtClosestPtPointTriangle(closest, p, a, b, c []float64) {
	// Check if P in vertex region outside A
	ab := [3]float64{}
	ac := [3]float64{}
	ap := [3]float64{}
	DtVsub(ab[:], b, a)
	DtVsub(ac[:], c, a)
	DtVsub(ap[:], p, a)
	d1 := DtVdot(ab[:], ap[:])
	d2 := DtVdot(ac[:], ap[:])
	if d1 <= 0.0 && d2 <= 0.0 {
		// barycentric coordinates (1,0,0)
		DtVcopy(closest, a)
		return
	}

	// Check if P in vertex region outside B
	bp := [3]float64{}
	DtVsub(bp[:], p, b)
	d3 := DtVdot(ab[:], bp[:])
	d4 := DtVdot(ac[:], bp[:])
	if d3 >= 0.0 && d4 <= d3 {
		// barycentric coordinates (0,1,0)
		DtVcopy(closest, b)
		return
	}

	// Check if P in edge region of AB, if so return projection of P onto AB
	vc := d1*d4 - d3*d2
	if vc <= 0.0 && d1 >= 0.0 && d3 <= 0.0 {
		// barycentric coordinates (1-v,v,0)
		v := d1 / (d1 - d3)
		closest[0] = a[0] + v*ab[0]
		closest[1] = a[1] + v*ab[1]
		closest[2] = a[2] + v*ab[2]
		return
	}

	// Check if P in vertex region outside C
	cp := [3]float64{}
	DtVsub(cp[:], p, c)
	d5 := DtVdot(ab[:], cp[:])
	d6 := DtVdot(ac[:], cp[:])
	if d6 >= 0.0 && d5 <= d6 {
		// barycentric coordinates (0,0,1)
		DtVcopy(closest, c)
		return
	}

	// Check if P in edge region of AC, if so return projection of P onto AC
	vb := d5*d2 - d1*d6
	if vb <= 0.0 && d2 >= 0.0 && d6 <= 0.0 {
		// barycentric coordinates (1-w,0,w)
		w := d2 / (d2 - d6)
		closest[0] = a[0] + w*ac[0]
		closest[1] = a[1] + w*ac[1]
		closest[2] = a[2] + w*ac[2]
		return
	}

	// Check if P in edge region of BC, if so return projection of P onto BC
	va := d3*d6 - d5*d4
	if va <= 0.0 && (d4-d3) >= 0.0 && (d5-d6) >= 0.0 {
		// barycentric coordinates (0,1-w,w)
		w := (d4 - d3) / ((d4 - d3) + (d5 - d6))
		closest[0] = b[0] + w*(c[0]-b[0])
		closest[1] = b[1] + w*(c[1]-b[1])
		closest[2] = b[2] + w*(c[2]-b[2])
		return
	}

	// P inside face region. Compute Q through its barycentric coordinates (u,v,w)
	denom := 1.0 / (va + vb + vc)
	v := vb * denom
	w := vc * denom
	closest[0] = a[0] + ab[0]*v + ac[0]*w
	closest[1] = a[1] + ab[1]*v + ac[1]*w
	closest[2] = a[2] + ab[2]*v + ac[2]*w
}

var EPS float64 = 1e-4

/// Derives the y-axis height of the closest point on the triangle from the specified reference point.
///  @param[in]		p		The reference point from which to test. [(x, y, z)]
///  @param[in]		a		Vertex A of triangle ABC. [(x, y, z)]
///  @param[in]		b		Vertex B of triangle ABC. [(x, y, z)]
///  @param[in]		c		Vertex C of triangle ABC. [(x, y, z)]
///  @param[out]	h		The resulting height.
func DtClosestHeightPointTriangle(p, a, b, c []float64, h *float64) bool {
	v0 := [3]float64{}
	v1 := [3]float64{}
	v2 := [3]float64{}
	DtVsub(v0[:], c, a)
	DtVsub(v1[:], b, a)
	DtVsub(v2[:], p, a)

	dot00 := DtVdot2D(v0[:], v0[:])
	dot01 := DtVdot2D(v0[:], v1[:])
	dot02 := DtVdot2D(v0[:], v2[:])
	dot11 := DtVdot2D(v1[:], v1[:])
	dot12 := DtVdot2D(v1[:], v2[:])

	// Compute barycentric coordinates
	invDenom := 1.0 / (dot00*dot11 - dot01*dot01)
	u := (dot11*dot02 - dot01*dot12) * invDenom
	v := (dot00*dot12 - dot01*dot02) * invDenom

	// The (sloppy) epsilon is needed to allow to get height of points which
	// are interpolated along the edges of the triangles.
	//	static const float EPS = 1e-4f;

	// If point lies inside the triangle, return interpolated ycoord.
	if u >= -EPS && v >= -EPS && (u+v) <= 1+EPS {
		*h = a[1] + v0[1]*u + v1[1]*v
		return true
	}

	return false
}

func DtIntersectSegmentPoly2D(p0, p1, verts []float64, nverts int, tmin, tmax *float64, segMin, segMax *int) bool {
	*tmin = 0
	*tmax = 1
	*segMin = -1
	*segMax = -1

	dir := [3]float64{}
	DtVsub(dir[:], p1, p0)

	for i, j := 0, nverts-1; i < nverts; j, i = i, i+1 {
		edge := [3]float64{}
		diff := [3]float64{}
		DtVsub(edge[:], verts[i*3:], verts[j*3:])
		DtVsub(diff[:], p0, verts[j*3:])
		n := DtVperp2D(edge[:], diff[:])
		d := DtVperp2D(dir[:], edge[:])
		if math.Abs(d) < 0.00000001 {
			// S is nearly parallel to this edge
			if n < 0 {
				return false
			} else {
				continue
			}
		}
		t := n / d
		if d < 0 {
			// segment S is entering across this edge
			if t > *tmin {
				*tmin = t
				*segMin = j
				// S enters after leaving polygon
				if *tmin > *tmax {
					return false
				}
			}
		} else {
			// segment S is leaving across this edge
			if t < *tmax {
				*tmax = t
				*segMax = j
				// S leaves before entering polygon
				if *tmax < *tmin {
					return false
				}
			}
		}
	}

	return true
}

func vperpXZ(a, b []float64) float64 {
	return a[0]*b[2] - a[2]*b[0]
}

func DtIntersectSegSeg2D(ap, aq, bp, bq []float64, s, t *float64) bool {
	u := [3]float64{}
	v := [3]float64{}
	w := [3]float64{}
	DtVsub(u[:], aq, ap)
	DtVsub(v[:], bq, bp)
	DtVsub(w[:], ap, bp)
	d := vperpXZ(u[:], v[:])
	if math.Abs(d) < 1e-6 {
		return false
	}
	*s = vperpXZ(v[:], w[:]) / d
	*t = vperpXZ(u[:], w[:]) / d
	return true
}

/// Determines if the specified point is inside the convex polygon on the xz-plane.
///  @param[in]		pt		The point to check. [(x, y, z)]
///  @param[in]		verts	The polygon vertices. [(x, y, z) * @p nverts]
///  @param[in]		nverts	The number of vertices. [Limit: >= 3]
/// @return True if the point is inside the polygon.
func DtPointInPolygon(pt, verts []float64, nverts int) bool {
	var i, j int
	c := false
	for i, j = 0, nverts-1; i < nverts; j, i = i, i+1 {
		vi := verts[i*3:]
		vj := verts[j*3:]
		if ((vi[2] > pt[2]) != (vj[2] > pt[2])) &&
			(pt[0] < (vj[0]-vi[0])*(pt[2]-vi[2])/(vj[2]-vi[2])+vi[0]) {
			c = !c
		}
	}
	return c
}

func DtDistancePtSegSqr2D(pt, p, q []float64, t *float64) float64 {
	pqx := q[0] - p[0]
	pqz := q[2] - p[2]
	dx := pt[0] - p[0]
	dz := pt[2] - p[2]
	d := pqx*pqx + pqz*pqz
	*t = pqx*dx + pqz*dz
	if d > 0 {
		*t /= d
	}
	if *t < 0 {
		*t = 0
	} else if *t > 1 {
		*t = 1
	}
	dx = p[0] + (*t)*pqx - pt[0]
	dz = p[2] + (*t)*pqz - pt[2]
	return dx*dx + dz*dz
}

func DtDistancePtPolyEdgesSqr(pt, verts []float64, nverts int, ed, et []float64) bool {
	var i, j int
	c := false
	for i, j = 0, nverts-1; i < nverts; j, i = i, i+1 {
		vi := verts[i*3:]
		vj := verts[j*3:]
		if ((vi[2] > pt[2]) != (vj[2] > pt[2])) &&
			(pt[0] < (vj[0]-vi[0])*(pt[2]-vi[2])/(vj[2]-vi[2])+vi[0]) {
			c = !c
		}
		ed[j] = DtDistancePtSegSqr2D(pt, vj, vi, &et[j])
	}
	return c
}

/// Derives the centroid of a convex polygon.
///  @param[out]	tc		The centroid of the polgyon. [(x, y, z)]
///  @param[in]		idx		The polygon indices. [(vertIndex) * @p nidx]
///  @param[in]		nidx	The number of indices in the polygon. [Limit: >= 3]
///  @param[in]		verts	The polygon vertices. [(x, y, z) * vertCount]
func DtCalcPolyCenter(tc []float64, idx []uint16, nidx int, verts []float64) {
	tc[0] = 0.0
	tc[1] = 0.0
	tc[2] = 0.0
	for j := 0; j < nidx; j++ {
		v := verts[idx[j]*3:]
		tc[0] += v[0]
		tc[1] += v[1]
		tc[2] += v[2]
	}
	s := 1.0 / float64(nidx)
	tc[0] *= s
	tc[1] *= s
	tc[2] *= s
}

func projectPoly(axis, poly []float64, npoly int, rmin, rmax *float64) {
	*rmax = DtVdot2D(axis, poly)
	*rmin = *rmax
	for i := 1; i < npoly; i++ {
		d := DtVdot2D(axis, poly[i*3:])
		*rmin = DtMinFloat64(*rmin, d)
		*rmax = DtMaxFloat64(*rmax, d)
	}
}

func overlapRange(amin, amax, bmin, bmax, eps float64) bool {
	return !((amin+eps) > bmax || (amax-eps) < bmin)
}

/// Determines if the two convex polygons overlap on the xz-plane.
///  @param[in]		polya		Polygon A vertices.	[(x, y, z) * @p npolya]
///  @param[in]		npolya		The number of vertices in polygon A.
///  @param[in]		polyb		Polygon B vertices.	[(x, y, z) * @p npolyb]
///  @param[in]		npolyb		The number of vertices in polygon B.
/// @return True if the two polygons overlap.
func DtOverlapPolyPoly2D(polya []float64, npolya int, polyb []float64, npolyb int) bool {
	for i, j := 0, npolya-1; i < npolya; j, i = i, i+1 {
		va := polya[j*3:]
		vb := polya[i*3:]
		n := [3]float64{vb[2] - va[2], 0, -(vb[0] - va[0])}
		var amin, amax, bmin, bmax float64
		projectPoly(n[:], polya, npolya, &amin, &amax)
		projectPoly(n[:], polyb, npolyb, &bmin, &bmax)
		if !overlapRange(amin, amax, bmin, bmax, EPS) {
			// Found separating axis
			return false
		}
	}
	for i, j := 0, npolyb-1; i < npolyb; j, i = i, i+1 {
		va := polyb[j*3:]
		vb := polyb[i*3:]
		n := [3]float64{vb[2] - va[2], 0, -(vb[0] - va[0])}
		var amin, amax, bmin, bmax float64
		projectPoly(n[:], polya, npolya, &amin, &amax)
		projectPoly(n[:], polyb, npolyb, &bmin, &bmax)
		if !overlapRange(amin, amax, bmin, bmax, EPS) {
			// Found separating axis
			return false
		}
	}
	return true
}
