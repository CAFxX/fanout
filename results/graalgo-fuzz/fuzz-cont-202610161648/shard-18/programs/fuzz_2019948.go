package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 int; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint16 }
func genFn0() int16 {
    v6 := []int{32768, 32767}
    switch 42 {
    case 127:
        var v7 int8 = 127
        _ = v7
    case 32768:
        if true {
            if len(v6) > 0 {
                v6[0] = -1
            }
            switch 4294967295 {
            case 4294967295:
                _ = len(v6)
            case -460:
                var v8 int32 = -677
                _ = v8
            case 795:
                var v9 uint8 = 14
                _ = v9
            default:
                for fi10 := 0; fi10 < 15; fi10++ {
                    var v11 uint64 = 2
                    _ = fi10
                    _ = v11
                }
            }
            var v12 uint = 42
            _ = v12
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                if false {
                    if false {
                        var v14 int64 = 515
                        _ = v14
                    } else {
                        if true {
                            if len(v6) > 0 {
                                v6[0] = (fi13 >> ((510) & 63))
                            }
                            switch "foo bar" {
                            case "hello":
                                var v15 int64 = 9223372036854775807
                                _ = v15
                            default:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        } else {
                            for fi16 := 0; fi16 < 8; fi16++ {
                                var v17 int8 = 127
                                _ = fi16
                                _ = v17
                            }
                            fmt.Println(fi13)
                        }
                    }
                }
                _ = fi13
            }
            var v18 uint64 = 0
            var v19 uint8 = 0
            _ = v18
            _ = v19
        }
    case -2147483648:
        if (175 <= 42) {
            var v20 uint32 = 4294967295
            switch "hello" {
            case "a":
                var v21 uint16 = 463
                _ = v21
            case "foo bar":
                v20 = (v20 * (v20 << ((2) & 63)))
            }
            _ = v20
        }
    }
    var v22 float32 = 1e10
    var v23 uint8 = 210
    _ = v22
    _ = v23
    _ = v6
    return 99
}
func genFn1(p0 bool, p1 bool) bool {
    var v24 uint = 255
    for fi25 := 0; fi25 < 20; fi25++ {
        if ((true && (!true)) && ((42 <= -1) && p0)) {
            v26 := []int{fi25, fi25}
            var v27 int64 = 0
            _ = v26
            _ = v27
        } else {
            var v28 uint8 = 127
            _ = v28
        }
        _ = fi25
    }
    switch -2147483648 {
    case 2147483647:
        p1 = true
    default:
        v29 := GenS1{F0: 65535, F1: 63}
        _ = v29
    }
    fmt.Println(p0)
    switch "hello" {
    case "a":
        var v30 uint64 = 1
        _ = v30
    case "hello":
        fmt.Println(v24, v24)
    default:
        v31 := []int{4294967295, 32768, 32768, 510}
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v24
    return p0
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{-76, -666, 0, 4294967295}
    v3 := make(map[string]bool)
    v4 := []int{-911, 9223372036854775807, 2147483647}
    var v5 float32 = v1
    for fi32 := 0; fi32 < 18; fi32++ {
        fmt.Println(v1, v5)
        for fi33 := 0; fi33 < 6; fi33++ {
            fmt.Println(fi32, fi32, v5)
            if len(v2) > 0 {
                v2[0] = fi33
            }
            _ = fi33
        }
        _ = fi32
    }
    var v34 int = 2147483647
    if false {
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 string = ("hello" + "\n\t")
            v37 := GenS2{F0: (v36 + ((v36 + (v36 + "hello")) + "foo bar")), F1: v34, F2: 0}
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    v38 := v3[""]
    _ = len(v3)
    fmt.Println(v38, v1)
    v39 := GenS1{F0: 109, F1: 42}
    v38 = false
    v40 := GenS3{F0: 0}
    v41 := GenS4{F0: 4294967295, F1: 518}
    v4 = append(v4, v34)
    v42 := GenS2{F0: "\n\t", F1: (-(v34)), F2: 1000}
    var v43 int = v34
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
