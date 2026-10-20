package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 int16; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int64; F2 int }
type GenS5 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint8, p1 int32, p2 uint) int64 {
    var v6 int16 = 0
    v7 := []int{2147483647, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 568
}
func genFn1(p0 int8, p1 uint8) int {
    if ((0.0 - 0.1) != 3.14159) {
        if (!(4294967295 == 255)) {
            p0 = (p0 << ((0) & 63))
            var v8 float64 = 0.1
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 bool = true
                if ((fi9 << ((4294967295) & 63)) == fi9) {
                    var v11 float32 = -2.25
                    v12 := GenS1{F0: 1, F1: 42, F2: (255 != 65535)}
                    _ = v11
                    _ = v12
                }
                var v13 int8 = (^(p0))
                _ = fi9
                _ = v10
                _ = v13
            }
            _ = v8
        }
        v14 := GenS2{F0: 744, F1: 21, F2: 2}
        _ = v14
    }
    var v15 float64 = (-2.25 - 0.0)
    var v16 int32 = 0
    if true {
        p1 = p1
        v15 = v15
        fmt.Println(v15, v15, p0)
    }
    p0 = -100
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 32767
}

func main() {
    var v1 float64 = ((1e10 + 3.14159) / 1.5)
    var v2 int = 42
    var v3 int32 = 1
    var v4 uint64 = 2
    var v5 int64 = -803
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 string = "hello"
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 float32 = 1.5
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 bool = true
            v23 := []int{}
            v22 = ((((v18 + v18) + v18) + "\n\t") != ((v18 + (v18 + "\n\t")) + v18))
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi17
        _ = v18
    }
    switch ("foo bar" + "foo bar") {
    case "":
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS3{F0: 453}
            _ = fi24
            _ = v25
        }
    default:
        var v26 int16 = 0
        _ = v26
    }
    v4 = v4
    if ((0 < 0) || (!true)) {
        fmt.Println(v2, v2)
    }
    var v27 uint8 = 1
    var v28 float32 = 3.14159
    if (!(952 < uint16(v4))) {
        if true {
            v29 := []int{v2, v2, 32768}
            _ = v29
        } else {
            v4 += uint64(v27)
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                v30 := GenS4{F0: v3, F1: v5, F2: v2}
                _ = v30
            case "hello":
                v31 := GenS5{F0: (v4 & 2), F1: 783}
                _ = v31
            case "":
                var v32 int32 = (v3 >> ((225) & 63))
                _ = v32
            }
            v33 := []int{v2, (-(v2)), 32767}
            _ = v33
        }
        v2 -= (v2 * (v2 >> ((255) & 63)))
        for fi34 := 0; fi34 < 20; fi34++ {
            fmt.Println(v4)
            if true {
                v35 := []int{v2, (fi34 | 846)}
                var v36 uint32 = 1
                var v37 float64 = v1
                _ = v35
                _ = v36
                _ = v37
            } else {
                var v38 uint64 = v4
                var v39 float64 = 0.1
                fmt.Println(v39)
                _ = v38
                _ = v39
            }
            switch "foo bar" {
            case "":
                var v40 int8 = 0
                _ = v40
            }
            _ = fi34
        }
    }
    v41 := []int{(v2 - v2), (v2 + (v2 << ((18446744073709551615) & 63))), -2147483648}
    v28 += v28
    if false {
        v5 = (v5 >> ((18446744073709551615) & 63))
        if true {
            v42 := GenS4{F0: v3, F1: (v5 % ((v5) | 1)), F2: 0}
            fmt.Println(v5)
            v3 = v3
            _ = v42
        } else {
            switch ("a" + "\n\t") {
            case "\n\t":
                for fi43 := 0; fi43 < 15; fi43++ {
                    v27 = (v27 * (v27 - (^(v27))))
                    _ = fi43
                }
            }
        }
    }
    fmt.Println(v27, v2)
    var v44 int = (v2 & v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v5
}
