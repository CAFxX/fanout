package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 bool }
func genFn0(p0 string, p1 int32) bool {
    p1 = p1
    var v6 float32 = (3.14159 + 0.0)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := []int{128, -2147483648}
    var v2 string = "a"
    var v3 float64 = 0.0
    v4 := make(map[string]uint32)
    var v5 uint = 732
    if (0 < 1) {
        var v8 string = v2
        v1 = append(v1, -601)
        switch v8 {
        case "a":
            fmt.Println(v5, v3, v2)
        case "":
            if (496 != 18446744073709551615) {
                v3 += 1.5
                v1 = append(v1, 42)
            }
        }
        _ = v8
    }
    fmt.Println(v2, v3)
    for fi9 := 0; fi9 < 10; fi9++ {
        v1 = append(v1, 127)
        var v10 int = (fi9 % ((fi9) | 1))
        _ = fi9
        _ = v10
    }
    v5 = v5
    var v11 uint16 = uint16(v5)
    var v12 uint = 0
    v13 := v4[v2]
    switch 9223372036854775807 {
    case 419:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 1:
        var v14 uint32 = (v13 - 0)
        _ = v14
    }
    v3 += (3.14159 + 1.5)
    var v15 bool = true
    for fi16 := 0; fi16 < 16; fi16++ {
        v15 = ("a" != v2)
        v1 = append(v1, 9223372036854775807)
        if ((-(fi16)) >= fi16) {
            v17 := v4["hello"]
            _ = v17
        } else {
            v18 := GenS1{F0: 861}
            if (32767 <= fi16) {
                fmt.Println(fi16, v2)
            }
            v19 := GenS1{F0: v11}
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    delete(v4, v2)
    var v20 uint64 = 1
    if (!(v11 < (^(v11)))) {
        for fi21 := 0; fi21 < 2; fi21++ {
            for fi22 := 0; fi22 < 1; fi22++ {
                switch v2 {
                case "a":
                    v23 := GenS1{F0: v11}
                    _ = v23
                case "foo bar":
                    _ = len(v4)
                default:
                    if (!((true && false) || v15)) {
                        var v24 uint32 = v13
                        v20 = (v20 - uint64(v13))
                        delete(v4, v2)
                        v24 += v13
                        _ = v24
                    } else {
                        switch v2 {
                        case "":
                            fmt.Println(fi22, fi22, v15)
                        }
                        delete(v4, (v2 + (v2 + ("" + (v2 + "hello")))))
                        switch (fi21 | fi21) {
                        case 1:
                            switch v2 {
                            case "hello":
                                v2 = (((v2 + v2) + v2) + v2)
                            case "foo bar":
                                fmt.Println(v15)
                            case "a":
                                if false {
                                    v25 := GenS1{F0: (^(v11))}
                                    var v26 uint64 = (v20 & 2)
                                    fmt.Println(v20)
                                    _ = v25
                                    _ = v26
                                } else {
                                    v11 = v11
                                    var v27 float32 = (1.5 - 3.14159)
                                    _ = v27
                                }
                            }
                        case 9223372036854775807:
                            var v28 float32 = 1e10
                            _ = v28
                        case -901:
                            v29 := GenS2{F0: fi21}
                            _ = v29
                        default:
                            switch fi22 {
                            case -2147483648:
                                var v30 uint64 = v20
                                _ = v30
                            case -9223372036854775808:
                                var v31 int64 = 42
                                _ = v31
                            case 1:
                                v32 := GenS2{F0: fi22}
                                _ = v32
                            }
                        }
                    }
                }
                v33 := GenS3{F0: 255, F1: (v13 << (((v12 | (v12 | v5))) & 63)), F2: (!v15)}
                _ = fi22
                _ = v33
            }
            var v34 float32 = (0.0 / float32(v3))
            for fi35 := 0; fi35 < 12; fi35++ {
                if v15 {
                    v12 -= 0
                    fmt.Println(v12, v13)
                } else {
                    v1 = append(v1, (fi35 << ((v5) & 63)))
                    v12 = v12
                    var v36 uint = v12
                    _ = v36
                }
                _ = fi35
            }
            _ = fi21
            _ = v34
        }
    } else {
        var v37 int64 = -335
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
