package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint16; F2 int16 }
func genFn0(p0 int, p1 int32) int {
    if (!true) {
        fmt.Println(p1)
    } else {
        p0 = p0
        p0 -= (^(p0))
    }
    v7 := []int{(-(p0)), p0, (^(p0)), 127}
    v8 := GenS1{F0: 52, F1: -1, F2: true}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p0 - p0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 0
    v3 := make(map[uint32]int)
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    var v6 int = 32768
    var v9 int64 = -707
    if true {
        var v10 int8 = 42
        v11 := v1[233]
        _ = v10
        _ = v11
    } else {
        var v12 float32 = 1e10
        for fi13 := 0; fi13 < 14; fi13++ {
            v2 += 18446744073709551615
            v6 -= genFn0(v6, 2147483647)
            v14 := v5[4294967295]
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v1[1] = 2147483647
    if ((!true) && (100 >= -75)) {
        if false {
            v15 := GenS1{F0: 42, F1: -1, F2: (0 != -1)}
            for fi16 := 0; fi16 < 17; fi16++ {
                switch (fi16 << ((4294967295) & 63)) {
                case 898:
                    switch (("\n\t" + ("\n\t" + ("hello" + ("foo bar" + "")))) + "") {
                    case "":
                        switch "a" {
                        case "":
                            v17 := v4[((("hello" + "hello") + "a") + "\n\t")]
                            _ = v17
                        case "\n\t":
                            v6 = (fi16 + (v6 * fi16))
                        default:
                            var v18 uint32 = 2
                            _ = v18
                        }
                    case "foo bar":
                        v6 = (v6 % ((int(v9)) | 1))
                    }
                case 959:
                    switch "\n\t" {
                    case "hello":
                        switch "hello" {
                        case "":
                            var v19 int = -775
                            _ = v19
                        case "foo bar":
                            var v20 bool = (!((((-2.25 + float32(3.14159)) - -2.25) - 0.1) >= (-2.25 - 0.1)))
                            _ = v20
                        }
                    case "foo bar":
                        v21 := []int{(-(fi16))}
                        _ = v21
                    }
                }
                v22 := GenS2{F0: 1, F1: (v2 / (((^(v2))) | 1)), F2: (v2 * v2)}
                for fi23 := 0; fi23 < 12; fi23++ {
                    if (1.5 <= 0.0) {
                        var v24 int = (fi23 - (fi16 >> ((42) & 63)))
                        var v25 uint8 = 2
                        _ = v24
                        _ = v25
                    } else {
                        v26 := []int{fi16}
                        _ = v26
                    }
                    if (false && false) {
                        var v27 int32 = -744
                        v28 := GenS2{F0: 32767, F1: 475, F2: 18446744073709551615}
                        if true {
                            var v29 float32 = 1e10
                            v29 -= v29
                            v30 := []int{int(v2)}
                            _ = v29
                            _ = v30
                        }
                        _ = v27
                        _ = v28
                    } else {
                        v31 := GenS3{F0: 100}
                        v9 += -9223372036854775808
                        v9 = v9
                        _ = v31
                    }
                    switch (fi16 / ((-505) | 1)) {
                    case 0:
                        fmt.Println(fi23, v2)
                    case 127:
                        v32 := GenS4{F0: "foo bar", F1: 255, F2: int16(v2)}
                        _ = v32
                    default:
                        var v33 int = fi16
                        _ = v33
                    }
                    var v34 uint64 = (v2 * v2)
                    _ = fi23
                    _ = v34
                }
                _ = fi16
                _ = v22
            }
            for fi35 := 0; fi35 < 18; fi35++ {
                fmt.Println(v9, v6)
                _ = fi35
            }
            _ = v15
        } else {
            v36 := []int{genFn0(v6, -827), 0, v6}
            v37 := GenS3{F0: 42}
            _ = v36
            _ = v37
        }
        var v38 uint8 = 0
        _ = v38
    }
    if (int(v2) != v6) {
        switch v6 {
        case 544:
            v39 := v4["\n\t"]
            _ = v39
        case 32768:
            fmt.Println(v6, v9, v9)
        case -469:
            var v40 uint = 2
            _ = v40
        }
    }
    if true {
        fmt.Println(v6, v9)
    }
    if false {
        v41 := GenS1{F0: 390, F1: -903, F2: (true && false)}
        v2 = (v2 | (v2 / ((v2) | 1)))
        _ = v41
    } else {
        v2 -= (v2 >> ((34) & 63))
        fmt.Println(v6, v9, v9)
    }
    switch 128 {
    case 127:
        for fi42 := 0; fi42 < 15; fi42++ {
            v43 := GenS1{F0: 65535, F1: 0, F2: false}
            var v44 string = "foo bar"
            v6 += (fi42 >> ((2) & 63))
            _ = fi42
            _ = v43
            _ = v44
        }
    case 32767:
        v6 -= -9223372036854775808
    }
    v2 = 1
    v45 := []int{(v6 & v6), (v6 >> ((65535) & 63)), -9223372036854775808}
    v46 := GenS2{F0: 1, F1: (v2 / ((v2) | 1)), F2: v2}
    if (0.1 != -2.25) {
        switch "a" {
        case "":
            if len(v45) > 0 {
                v45[0] = v6
            }
        default:
            for fi47 := 0; fi47 < 14; fi47++ {
                v6 = (fi47 ^ genFn0(42, -2147483648))
                var v48 uint16 = 32767
                v9 += v9
                _ = fi47
                _ = v48
            }
        }
        var v49 int16 = -129
        _ = v49
    }
    for fi50 := 0; fi50 < 3; fi50++ {
        switch v6 {
        case 128:
            if true {
                fmt.Println(v6)
                var v51 bool = (!true)
                _ = v51
            } else {
                var v52 uint = 18446744073709551615
                var v53 int8 = 127
                _ = v52
                _ = v53
            }
        }
        _ = fi50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
