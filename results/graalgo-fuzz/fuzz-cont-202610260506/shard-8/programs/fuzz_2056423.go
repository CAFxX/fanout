package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
func genFn0() int64 {
    v7 := GenS1{F0: 18}
    if (!(!false)) {
        var v8 bool = (42 != 282)
        if v8 {
            v9 := []int{}
            _ = v9
        } else {
            v10 := GenS1{F0: 40}
            _ = v10
        }
        fmt.Println(v8)
        _ = v8
    } else {
        var v11 uint = 42
        _ = v11
    }
    var v12 string = ("a" + (("" + "\n\t") + "\n\t"))
    var v13 string = (v12 + ("" + (v12 + v12)))
    switch 0 {
    case -459:
        var v14 int32 = 42
        _ = v14
    case 0:
        v15 := []int{769}
        _ = v15
    case 32767:
        if true {
            v12 = "hello"
            var v16 int16 = -516
            _ = v16
        }
    }
    _ = v12
    _ = v13
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 int, p1 uint64, p2 uint64) uint32 {
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := []int{}
        switch fi17 {
        case -488:
            switch 9223372036854775807 {
            case 175:
                v19 := GenS2{F0: 2}
                _ = v19
            default:
                var v20 string = "a"
                _ = v20
            }
        case 170:
            v18 = append(v18, -2147483648)
        default:
            if len(v18) > 0 {
                v18[0] = fi17
            }
        }
        p0 += fi17
        _ = fi17
        _ = v18
    }
    fmt.Println(p0, p2, p2)
    if (127 == 42) {
        if (!true) {
            p2 = 42
            for fi21 := 0; fi21 < 6; fi21++ {
                p0 += 32768
                v22 := []int{fi21, 805, -768, (fi21 * p0)}
                var v23 int64 = 1
                _ = fi21
                _ = v22
                _ = v23
            }
        } else {
            var v24 float32 = ((-2.25 + -2.25) + float32(0.0))
            for fi25 := 0; fi25 < 9; fi25++ {
                fmt.Println(p0)
                var v26 uint64 = 4294967295
                v27 := GenS1{F0: 128}
                _ = fi25
                _ = v26
                _ = v27
            }
            v28 := []int{p0, (p0 % ((p0) | 1)), p0, (p0 ^ p0)}
            _ = v24
            _ = v28
        }
        v29 := []int{}
        p1 = p2
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 int64 = -835
    v2 := []int{0, -1, 127, -67}
    var v3 int16 = 42
    var v4 uint32 = 0
    var v5 float32 = 0.0
    var v6 uint64 = 18446744073709551615
    v1 -= (v1 << ((18446744073709551615) & 63))
    if true {
        for fi30 := 0; fi30 < 8; fi30++ {
            if false {
                for fi31 := 0; fi31 < 17; fi31++ {
                    v6 = (v6 % (((v6 & (v6 - 2))) | 1))
                    _ = fi31
                }
                fmt.Println(fi30)
            } else {
                var v32 uint = 42
                var v33 bool = (!(1e10 != v5))
                _ = v32
                _ = v33
            }
            v34 := GenS3{F0: (^(v4))}
            v35 := GenS3{F0: (v4 / ((2147483648) | 1))}
            _ = fi30
            _ = v34
            _ = v35
        }
    }
    v2 = append(v2, -1)
    var v36 float32 = 1e10
    v5 += 3.14159
    var v37 uint = 18446744073709551615
    if false {
        if false {
            var v38 int64 = (-(v1))
            _ = v38
        }
        v6 = 4294967295
        var v39 uint16 = 1
        _ = v39
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        _ = len(v2)
        _ = fi40
    }
    v6 -= 159
    for fi41 := 0; fi41 < 5; fi41++ {
        for fi42 := 0; fi42 < 14; fi42++ {
            if len(v2) > 0 {
                v2[0] = (fi41 << ((18446744073709551615) & 63))
            }
            var v43 uint8 = 42
            switch 81 {
            case 2147483647:
                var v44 int32 = -624
                _ = v44
            }
            _ = fi42
            _ = v43
        }
        if true {
            fmt.Println(v4, fi41, fi41)
            _ = len(v2)
        }
        fmt.Println(v1, v1)
        _ = fi41
    }
    var v45 string = ("\n\t" + ("\n\t" + "\n\t"))
    v46 := GenS3{F0: (v4 & v4)}
    v47 := GenS4{F0: int32(v4)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v48 uint = (v37 - (v37 | v37))
    var v49 int64 = 42
    for fi50 := 0; fi50 < 19; fi50++ {
        fmt.Println(fi50, v36)
        for fi51 := 0; fi51 < 7; fi51++ {
            v36 = -2.25
            _ = fi51
        }
        _ = fi50
    }
    for fi52 := 0; fi52 < 4; fi52++ {
        _ = len(v2)
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
