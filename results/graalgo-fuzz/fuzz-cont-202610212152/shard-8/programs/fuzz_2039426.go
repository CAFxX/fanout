package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 string }
type GenS3 struct { F0 bool; F1 string; F2 bool }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint8) int {
    switch "\n\t" {
    case "":
        var v7 int64 = -1
        _ = v7
    case "hello":
        var v8 int32 = 367
        _ = v8
    case "\n\t":
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    var v10 float32 = (0.0 / 1e10)
    var v11 string = (("" + "") + "a")
    if ((false && true) || false) {
        fmt.Println(p0)
        switch v11 {
        case "foo bar":
            v12 := []int{0, -9223372036854775808, 2147483647, 753}
            _ = v12
        case "hello":
            for fi13 := 0; fi13 < 8; fi13++ {
                switch (fi13 / (((fi13 / (((fi13 ^ fi13)) | 1))) | 1)) {
                case -165:
                    var v14 int = fi13
                    _ = v14
                case 81:
                    switch fi13 {
                    case -9223372036854775808:
                        switch (fi13 >> ((42) & 63)) {
                        case 42:
                            var v15 uint8 = p0
                            _ = v15
                        default:
                            if (1 > 737) {
                                for fi16 := 0; fi16 < 4; fi16++ {
                                    var v17 uint64 = 42
                                    _ = fi16
                                    _ = v17
                                }
                            } else {
                                switch v11 {
                                case "":
                                    var v18 int = -110
                                    _ = v18
                                case "foo bar":
                                    fmt.Println(p0, v10)
                                }
                                var v19 string = v11
                                _ = v19
                            }
                        }
                    case -117:
                        var v20 uint64 = 42
                        _ = v20
                    default:
                        v11 = ("\n\t" + (("" + v11) + "hello"))
                    }
                case 9223372036854775807:
                    v21 := []int{(fi13 * -323), 1, (^(fi13)), fi13}
                    _ = v21
                }
                v22 := GenS1{F0: (("hello" + v11) + v11), F1: p0, F2: -9223372036854775808}
                _ = fi13
                _ = v22
            }
        default:
            switch (v11 + (v11 + v11)) {
            case "\n\t":
                fmt.Println(v10)
            case "":
                for fi23 := 0; fi23 < 3; fi23++ {
                    fmt.Println(p0, fi23)
                    v24 := []int{}
                    _ = fi23
                    _ = v24
                }
            case "hello":
                var v25 uint16 = 2
                _ = v25
            }
        }
    } else {
        v11 = (("hello" + v11) + "\n\t")
    }
    _ = p0
    _ = v10
    _ = v11
    return -1
}
func genFn1(p0 int64, p1 bool) bool {
    var v26 string = ""
    switch 32767 {
    case 0:
        v27 := GenS2{F0: 1, F1: 0, F2: (("foo bar" + v26) + v26)}
        _ = v27
    }
    v28 := []int{32768, 127, -9223372036854775808, 4294967295}
    if (-32768 >= 32767) {
        if (!p1) {
            var v29 float64 = -2.25
            if p1 {
                var v30 uint8 = 0
                if (-2147483648 != 42) {
                    v30 += 1
                }
                var v31 float32 = 3.14159
                _ = v30
                _ = v31
            } else {
                v32 := GenS3{F0: p1, F1: "\n\t", F2: false}
                _ = v32
            }
            switch 42 {
            case 1:
                if false {
                    var v33 int32 = 0
                    var v34 int32 = (v33 & 42)
                    _ = v33
                    _ = v34
                } else {
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                }
            }
            _ = v29
        } else {
            v26 = v26
            var v35 int8 = -1
            _ = v35
        }
        v28 = append(v28, genFn0(2))
    } else {
        v36 := GenS2{F0: -2147483648, F1: 55, F2: v26}
        if p1 {
            if p1 {
                p1 = p1
                if (p0 == -9223372036854775808) {
                    switch v26 {
                    case "a":
                        v37 := GenS2{F0: -701, F1: -1, F2: "foo bar"}
                        _ = v37
                    }
                    switch "foo bar" {
                    case "":
                        switch v26 {
                        case "":
                            var v38 int16 = 32767
                            _ = v38
                        }
                    case "hello":
                        var v39 int16 = -769
                        _ = v39
                    case "foo bar":
                        v40 := GenS4{F0: 201}
                        _ = v40
                    }
                    v41 := GenS4{F0: 42}
                    _ = v41
                }
                for fi42 := 0; fi42 < 10; fi42++ {
                    for fi43 := 0; fi43 < 8; fi43++ {
                        fmt.Println(fi42)
                        var v44 uint8 = 128
                        _ = fi43
                        _ = v44
                    }
                    _ = fi42
                }
            } else {
                var v45 uint = 95
                _ = v45
            }
        }
        for fi46 := 0; fi46 < 3; fi46++ {
            v47 := GenS2{F0: 1, F1: 127, F2: v26}
            _ = fi46
            _ = v47
        }
        _ = v36
    }
    _ = p0
    _ = p1
    _ = v26
    _ = v28
    return p1
}

func main() {
    var v1 string = ("hello" + "\n\t")
    v2 := []int{-1, 776, 1}
    var v3 int64 = -9223372036854775808
    var v4 uint32 = 200
    var v5 uint = 189
    var v6 int32 = int32(v4)
    v6 = (v6 << ((1) & 63))
    switch 872 {
    case -591:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v48 := GenS2{F0: (^(v6)), F1: 42, F2: "hello"}
        _ = v48
    }
    var v49 int32 = (-(v6))
    for fi50 := 0; fi50 < 2; fi50++ {
        var v51 uint64 = 4294967295
        for fi52 := 0; fi52 < 12; fi52++ {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi52
        }
        fmt.Println(v1, v49)
        _ = fi50
        _ = v51
    }
    var v53 float64 = 3.14159
    switch 1 {
    case 2147483647:
        var v54 float32 = (float32(0.1) + 1.5)
        _ = v54
    case 127:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    _ = len(v2)
    var v55 uint64 = 4294967295
    var v56 int8 = -86
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v53
    _ = v55
    _ = v56
    _ = v6
}
