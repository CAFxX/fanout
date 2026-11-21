package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int }
func genFn0(p0 uint16) int32 {
    v5 := []int{474}
    if true {
        v6 := GenS1{F0: 0, F1: 800}
        p0 = p0
        _ = v6
    }
    switch 128 {
    case 0:
        v5 = append(v5, -414)
    case 9223372036854775807:
        p0 += 42
    }
    _ = p0
    _ = v5
    return -853
}
func genFn1() uint32 {
    v7 := GenS2{F0: 42}
    switch ("" + ("\n\t" + "")) {
    case "foo bar":
        var v8 uint32 = 65535
        _ = v8
    }
    v9 := GenS3{F0: 1, F1: 0}
    var v10 int32 = 340
    v11 := []int{}
    _ = v10
    _ = v11
    _ = v7
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 9223372036854775807
    var v3 string = "\n\t"
    var v4 float64 = (3.14159 / 3.14159)
    switch 813 {
    case 116:
        var v12 uint64 = 2
        _ = v12
    }
    fmt.Println(v4)
    for fi13 := 0; fi13 < 7; fi13++ {
        switch fi13 {
        case -1:
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int64 = (v2 << ((1) & 63))
                var v16 string = (v3 + (v3 + "hello"))
                if true {
                    v4 = v4
                    var v17 int64 = 9223372036854775807
                    v18 := GenS4{F0: 232, F1: 338}
                    _ = v17
                    _ = v18
                } else {
                    if ((((v3 + "a") + "hello") <= ("a" + v3)) && true) {
                        var v19 uint64 = 4294967295
                        v20 := []int{fi13}
                        v4 = 1e10
                        _ = v19
                        _ = v20
                    }
                    v21 := []int{4294967295, fi14}
                    _ = v21
                }
                var v22 uint = 726
                delete(v1, (v3 + "\n\t"))
                var v23 uint = (v22 % (((^(v22))) | 1))
                v22 -= v22
                _ = fi14
                _ = v15
                _ = v16
                _ = v22
                _ = v23
            }
        case -2147483648:
            v2 += (v2 & v2)
        default:
            if true {
                var v24 int8 = 0
                _ = v24
            }
        }
        if false {
            if (42 > -17) {
                var v25 uint = 255
                var v26 uint8 = 14
                fmt.Println(v26)
                _ = v25
                _ = v26
            } else {
                v27 := []int{42, fi13, 127}
                _ = v27
            }
        }
        if true {
            var v28 uint8 = 128
            v28 += (v28 / (((v28 / ((v28) | 1))) | 1))
            var v29 uint32 = 101
            _ = v28
            _ = v29
        }
        _ = fi13
    }
    var v30 int32 = genFn0(755)
    for fi31 := 0; fi31 < 8; fi31++ {
        switch fi31 {
        case 32767:
            switch (fi31 % ((fi31) | 1)) {
            case -2147483648:
                fmt.Println(v30, v4, fi31)
            default:
                var v32 string = v3
                _ = v32
            }
        default:
            var v33 uint = 255
            _ = v33
        }
        switch (^(fi31)) {
        case 128:
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 int16 = -1
                _ = fi34
                _ = v35
            }
        case -854:
            v3 = v3
        case 1:
            var v36 float64 = v4
            _ = v36
        default:
            v37 := GenS5{F0: 535}
            _ = v37
        }
        var v38 int16 = -1
        _ = fi31
        _ = v38
    }
    switch v3 {
    case "\n\t":
        var v39 int64 = (v2 & v2)
        _ = v39
    }
    v40 := GenS6{F0: 0}
    if true {
        if true {
            v3 = "hello"
            v41 := GenS2{F0: v2}
            _ = len(v1)
            _ = v41
        } else {
            for fi42 := 0; fi42 < 17; fi42++ {
                v2 = v2
                var v43 int8 = 1
                _ = fi42
                _ = v43
            }
        }
        v3 = ""
        for fi44 := 0; fi44 < 11; fi44++ {
            v45 := GenS7{F0: (fi44 * (fi44 % ((fi44) | 1)))}
            _ = fi44
            _ = v45
        }
    }
    v2 = 1
    v46 := []int{}
    switch v3 {
    case "a":
        v47 := GenS1{F0: 42, F1: -176}
        _ = v47
    case "hello":
        var v48 float64 = 1.5
        _ = v48
    case "":
        fmt.Println(v3, v30, v3)
    }
    var v49 int64 = -59
    var v50 uint = 599
    if len(v46) >= 2 {
        v46 = v46[1:len(v46)]
    }
    var v51 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v46
    _ = v49
    _ = v50
    _ = v51
}
