package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 int) uint32 {
    var v4 uint32 = 0
    if true {
        var v5 uint = 42
        v5 += (^(v5))
        switch (p0 | p0) {
        case 42:
            v6 := GenS1{F0: v5, F1: "foo bar", F2: -9223372036854775808}
            _ = v6
        case -199:
            var v7 uint64 = 3
            _ = v7
        }
        _ = v5
    } else {
        fmt.Println(p0, v4)
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
    }
    var v10 uint32 = (^(v4))
    var v11 int64 = -18
    p0 = 929
    _ = p0
    _ = v10
    _ = v11
    _ = v4
    return v10
}
func genFn1(p0 uint, p1 int8, p2 int64) uint64 {
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := []int{(fi12 / ((-9223372036854775808) | 1)), -1, (fi12 >> ((p0) & 63)), (fi12 % ((-9223372036854775808) | 1))}
        _ = len(v13)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if (((fi14 | (fi14 + 32768)) != -627) && true) {
            var v15 int8 = (^(p1))
            v16 := []int{fi14}
            fmt.Println(v15, fi14, p1)
            _ = v15
            _ = v16
        } else {
            var v17 uint32 = 1
            _ = v17
        }
        v18 := []int{fi14, (^(fi14))}
        p1 += 42
        _ = fi14
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 1
    var v2 int32 = int32(v1)
    var v3 int8 = 100
    if (false || true) {
        var v19 uint16 = (^(v1))
        for fi20 := 0; fi20 < 17; fi20++ {
            if false {
                var v21 uint64 = 2
                _ = v21
            } else {
                fmt.Println(fi20)
                fmt.Println(v19, v19, fi20)
                var v22 float64 = ((-2.25 * 0.0) * (0.1 / float64((-2.25 + 3.14159))))
                _ = v22
            }
            _ = fi20
        }
        switch ("" + "foo bar") {
        case "\n\t":
            if false {
                v3 = v3
            }
        case "hello":
            v3 += 42
        }
        _ = v19
    } else {
        v2 -= -1
        for fi23 := 0; fi23 < 18; fi23++ {
            v24 := []int{fi23, (fi23 + -2147483648)}
            _ = fi23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 uint8 = 127
        var v27 uint64 = 360
        var v28 float64 = 0.0
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    v29 := GenS2{F0: "hello"}
    v1 += 844
    v1 -= (v1 / (((v1 >> ((0) & 63))) | 1))
    var v30 int64 = 1
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := GenS2{F0: ("hello" + ("" + "a"))}
        _ = fi31
        _ = v32
    }
    v1 = v1
    if (!true) {
        var v33 uint64 = 2
        _ = v33
    } else {
        var v34 uint = 1000
        _ = v34
    }
    v2 += (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))
    if true {
        v2 -= -2147483648
        for fi35 := 0; fi35 < 12; fi35++ {
            switch "foo bar" {
            case "hello":
                var v36 int32 = v2
                _ = v36
            default:
                v37 := []int{fi35, (fi35 & (fi35 | fi35)), -279}
                _ = v37
            }
            _ = fi35
        }
    } else {
        v38 := GenS2{F0: ("a" + "")}
        _ = v38
    }
    var v39 uint32 = 42
    var v40 int64 = (v30 >> ((1) & 63))
    v41 := GenS3{F0: v3, F1: 1}
    v42 := []int{1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
