package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 int8) int64 {
    fmt.Println(p1, p1)
    var v6 string = ((("" + "foo bar") + "\n\t") + ("foo bar" + ""))
    _ = p0
    _ = p1
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 int16, p1 bool, p2 int16) int64 {
    var v7 int16 = (p2 >> ((18446744073709551615) & 63))
    if true {
        v8 := []int{464, 9223372036854775807, 32768, 32767}
        if true {
            for fi9 := 0; fi9 < 9; fi9++ {
                if p1 {
                    v10 := GenS1{F0: 2}
                    _ = v10
                }
                v7 = (-(p2))
                _ = fi9
            }
        }
        var v11 uint = 905
        _ = v11
        _ = v8
    } else {
        if true {
            if true {
                var v12 string = ("a" + "\n\t")
                _ = v12
            }
        }
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return genFn0(0, 53)
}

func main() {
    var v1 bool = false
    var v2 int = 730
    var v3 float32 = 1e10
    var v4 uint16 = 1
    v5 := make(map[string]uint32)
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS1{F0: 2}
        _ = fi13
        _ = v14
    }
    v2 = -9223372036854775808
    v15 := v5[("" + "foo bar")]
    var v16 uint64 = 0
    if (32767 <= v4) {
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := []int{v2, v2}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 float64 = float64((0.1 + 1e10))
            v21 := []int{414, (v2 - v2), -17, (fi19 >> ((2) & 63))}
            _ = fi19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int64 = -1
            _ = fi22
            _ = v23
        }
    } else {
        fmt.Println(v1)
        v15 -= (v15 / (((v15 / (((v15 * 2147483647)) | 1))) | 1))
    }
    if (!true) {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 int16 = 0
            if v1 {
                for fi26 := 0; fi26 < 7; fi26++ {
                    v1 = v1
                    if (int16(v16) <= 737) {
                        fmt.Println(v2, v2)
                        v25 = (v25 & (v25 << ((42) & 63)))
                        v27 := GenS2{F0: (!v1)}
                        v28 := GenS1{F0: 779}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 int = 1
                        _ = v29
                    }
                    v30 := GenS2{F0: (127 <= uint8(v15))}
                    var v31 int16 = (v25 % ((v25) | 1))
                    switch (("" + "hello") + "foo bar") {
                    case "a":
                        v1 = (v1 && (0.1 > 0.0))
                    case "foo bar":
                        var v32 float64 = ((0.0 * float64(0.0)) / 1e10)
                        _ = v32
                    }
                    _ = fi26
                    _ = v30
                    _ = v31
                }
                fmt.Println(v1)
            }
            v33 := GenS3{F0: v25}
            _ = fi24
            _ = v25
            _ = v33
        }
        var v34 int = (v2 << ((2) & 63))
        v2 += 32767
        _ = v34
    }
    v1 = ("hello" > "a")
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
