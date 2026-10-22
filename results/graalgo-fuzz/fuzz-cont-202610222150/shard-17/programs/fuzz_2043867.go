package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 string; F2 int }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int64 }
func genFn0(p0 string, p1 uint16, p2 int16) uint16 {
    p2 = (p2 >> ((42) & 63))
    var v6 uint8 = 2
    var v7 string = ("a" + "hello")
    if false {
        v8 := []int{4294967295, 128}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 int64, p2 uint16) int64 {
    if true {
        if (32767 >= 962) {
            if (118 == 889) {
                if ((42 >= (p1 * (p1 >> ((p0) & 63)))) && (32768 != 210)) {
                    p1 = (p1 % (((p1 ^ p1)) | 1))
                    for fi9 := 0; fi9 < 13; fi9++ {
                        var v10 uint64 = 4294967295
                        p2 = genFn0("\n\t", 32767, 1)
                        v11 := GenS1{F0: p0, F1: -1}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    p0 = 18446744073709551615
                }
                v12 := GenS1{F0: p0, F1: 77}
                _ = v12
            }
        }
        p1 += p1
        var v13 int = -2147483648
        _ = v13
    } else {
        v14 := GenS2{F0: 42}
        _ = v14
    }
    var v15 int64 = 42
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return v15
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "hello"
    var v3 string = v2
    v4 := make(map[uint32]int)
    var v5 bool = false
    v5 = (18446744073709551615 == 65535)
    var v16 int = -9223372036854775808
    if (v3 < v2) {
        var v17 int64 = 403
        _ = v17
    } else {
        var v18 int32 = 840
        _ = v18
    }
    if v5 {
        v3 = "foo bar"
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            v1[-825] = 2
            _ = fi19
        }
        v20 := GenS3{F0: 2, F1: (v3 + v3), F2: v16}
        switch ("foo bar" + v3) {
        case "":
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 float32 = float32(float64(1.5))
                var v23 uint = 42
                _ = fi21
                _ = v22
                _ = v23
            }
        case "a":
            v24 := []int{-2147483648, v16, v16}
            _ = v24
        default:
            if (!v5) {
                var v25 int8 = -91
                _ = v25
            } else {
                v26 := []int{v16, (v16 >> ((0) & 63))}
                _ = v26
            }
        }
        _ = v20
    }
    var v27 int = (^(v16))
    for fi28 := 0; fi28 < 13; fi28++ {
        v16 = v16
        switch v3 {
        case "foo bar":
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 float32 = ((((1.5 / 3.14159) / 0.1) + -2.25) + 0.1)
                _ = fi29
                _ = v30
            }
        case "":
            v31 := GenS4{F0: v16, F1: 65535}
            _ = v31
        case "a":
            var v32 uint32 = 2147483647
            _ = v32
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            v34 := GenS2{F0: 100}
            var v35 uint16 = genFn0(v2, 32767, -1)
            v35 = 272
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi28
    }
    var v36 float64 = 1e10
    if (!(196 > 226)) {
        var v37 int64 = 0
        v38 := v1[127]
        _ = v37
        _ = v38
    } else {
        v16 += 9223372036854775807
        if v5 {
            v39 := GenS3{F0: 1, F1: v3, F2: v16}
            fmt.Println(v27, v27)
            v40 := GenS5{F0: 656, F1: (v2 == v2)}
            _ = v39
            _ = v40
        }
        var v41 int8 = -8
        _ = v41
    }
    v2 = "a"
    v2 = ("" + ((v2 + v3) + (((v2 + v3) + (v3 + v2)) + "hello")))
    if true {
        v1[v16] = 65535
    }
    switch v2 {
    case "":
        fmt.Println(v2, v36)
    case "hello":
        v42 := v4[145]
        _ = v42
    }
    v43 := GenS6{F0: 844, F1: 775}
    switch (v27 * v16) {
    case 128:
        v44 := v4[1]
        _ = v44
    default:
        v4[1] = v27
    }
    switch v27 {
    case 32767:
        for fi45 := 0; fi45 < 11; fi45++ {
            v46 := []int{}
            _ = fi45
            _ = v46
        }
    case 204:
        v36 = float64(-2.25)
    default:
        v2 = ("a" + ("" + v3))
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        v48 := GenS7{F0: 845}
        for fi49 := 0; fi49 < 12; fi49++ {
            var v50 int64 = 42
            switch v27 {
            case 1:
                switch fi49 {
                case 1:
                    var v51 int32 = 1
                    _ = v51
                }
            default:
                var v52 bool = (23 <= 4294967295)
                _ = v52
            }
            if v5 {
                v50 = v50
            }
            _ = fi49
            _ = v50
        }
        _ = fi47
        _ = v48
    }
    var v53 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v43
    _ = v5
    _ = v53
}
