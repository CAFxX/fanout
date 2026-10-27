package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0(p0 int8) int32 {
    if (347 == -775) {
        var v6 string = "a"
        if (-206 == 42) {
            if true {
                var v7 int32 = -628
                v6 = ("\n\t" + "")
                if (v7 >= (v7 << ((255) & 63))) {
                    var v8 uint64 = 237
                    v9 := []int{406, -859, -379, -908}
                    _ = v8
                    _ = v9
                } else {
                    v10 := []int{42, -9223372036854775808, 32767}
                    _ = len(v10)
                    p0 = p0
                    _ = v10
                }
                _ = v7
            }
            if ((int16(p0) > 936) || false) {
                v11 := GenS1{F0: 65535}
                var v12 uint64 = 1
                _ = v11
                _ = v12
            } else {
                p0 += (p0 * 127)
                v13 := GenS1{F0: 1}
                v14 := []int{32767}
                _ = v13
                _ = v14
            }
        }
        v15 := []int{}
        _ = v15
        _ = v6
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        p0 -= 127
        _ = fi16
    }
    switch ("hello" + ("a" + "a")) {
    case "foo bar":
        for fi17 := 0; fi17 < 7; fi17++ {
            if ((42 < 541) || (0 != 389)) {
                v18 := GenS1{F0: 70}
                for fi19 := 0; fi19 < 3; fi19++ {
                    fmt.Println(fi19, fi17, fi19)
                    v20 := GenS1{F0: 42}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
            v21 := GenS2{F0: 827}
            _ = fi17
            _ = v21
        }
    case "hello":
        v22 := GenS2{F0: 63}
        _ = v22
    default:
        v23 := GenS3{F0: 65535, F1: 42}
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int32 = 435
        _ = fi24
        _ = v25
    }
    var v26 float64 = (0.0 * (float64(((3.14159 - 0.0) * 3.14159)) - (1.5 * (0.0 / 1.0))))
    _ = p0
    _ = v26
    return 42
}
func genFn1(p0 uint, p1 uint) uint32 {
    var v27 uint = p1
    var v28 uint = 4294967295
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 uint32 = 4294967295
        _ = fi29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return 2
}

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{1, 533}
    v3 := make(map[string]int)
    var v4 int16 = 1
    var v5 float64 = 3.14159
    for fi31 := 0; fi31 < 6; fi31++ {
        v3[""] = (fi31 & fi31)
        if (fi31 >= (fi31 | -355)) {
            v32 := GenS3{F0: 412, F1: 18446744073709551615}
            _ = v32
        } else {
            v2 = append(v2, 0)
            var v33 uint16 = v1
            _ = v33
        }
        var v34 string = ("\n\t" + ("hello" + ""))
        _ = fi31
        _ = v34
    }
    var v35 uint64 = 0
    var v36 int16 = (v4 - (v4 << ((656) & 63)))
    fmt.Println(v5, v5)
    for fi37 := 0; fi37 < 7; fi37++ {
        v38 := GenS1{F0: 378}
        _ = fi37
        _ = v38
    }
    var v39 uint32 = 42
    fmt.Println(v1, v1, v4)
    v36 -= v36
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
