package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int; F2 int32 }
type GenS6 struct { F0 uint; F1 int }
func genFn0(p0 uint64, p1 int8, p2 uint8) uint16 {
    var v7 uint32 = 2147483647
    p2 = (p2 >> ((65535) & 63))
    v8 := GenS1{F0: ((!true) && false), F1: uint(v7), F2: -2147483648}
    var v9 uint8 = (p2 % ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 string = ("foo bar" + ("" + "hello"))
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 uint = 746
    var v6 int64 = 1
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int = 32768
        switch -694 {
        case 2147483647:
            v3[v1] = ("" + "hello")
        }
        v11 -= fi10
        _ = fi10
        _ = v11
    }
    v5 = (v5 >> ((42) & 63))
    var v12 float64 = float64(float32(0.1))
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: false, F1: (v5 + (v5 + (v5 & (v5 % (((v5 << (((^(v5))) & 63))) | 1))))), F2: 42}
        _ = fi13
        _ = v14
    }
    fmt.Println(v6)
    fmt.Println(v5, v12)
    if (-1 >= -32768) {
        v15 := []int{}
        switch "a" {
        case "a":
            v16 := v3["\n\t"]
            _ = v16
        case "foo bar":
            switch "\n\t" {
            case "hello":
                for fi17 := 0; fi17 < 6; fi17++ {
                    if false {
                        v18 := GenS2{F0: 255, F1: -32768}
                        _ = v18
                    } else {
                        v19 := GenS3{F0: 42, F1: -1}
                        fmt.Println(v12)
                        v5 -= (v5 - v5)
                        _ = v19
                    }
                    var v20 uint32 = 2
                    _ = fi17
                    _ = v20
                }
            case "":
                if true {
                    v6 += -930
                    if len(v15) > 0 {
                        v15[0] = -2147483648
                    }
                }
            case "\n\t":
                _ = len(v3)
            }
        case "\n\t":
            for fi21 := 0; fi21 < 18; fi21++ {
                switch (fi21 ^ 128) {
                case 32768:
                    v15 = append(v15, (fi21 - (fi21 & (fi21 / ((165) | 1)))))
                case 128:
                    for fi22 := 0; fi22 < 17; fi22++ {
                        if true {
                            if (!(2147483647 <= 2147483648)) {
                                if true {
                                    fmt.Println(fi21, v5)
                                    if false {
                                        _ = len(v15)
                                        v6 = (^(v6))
                                        fmt.Println(v12, fi21, v12)
                                    } else {
                                        v23 := GenS2{F0: 2, F1: 1000}
                                        v24 := GenS4{F0: 0, F1: 0, F2: 4294967295}
                                        var v25 bool = false
                                        v5 += v5
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 uint32 = 1
                                    switch fi22 {
                                    case 32767:
                                        v27 := GenS1{F0: ((42 <= 127) && false), F1: (^(v5)), F2: -1}
                                        _ = v27
                                    case -9223372036854775808:
                                        switch "\n\t" {
                                        case "foo bar":
                                            v28 := GenS1{F0: false, F1: v5, F2: 883}
                                            _ = v28
                                        case "hello":
                                            v12 = 3.14159
                                        }
                                        v1 = v1
                                        v29 := GenS4{F0: -16, F1: (^(v26)), F2: 65535}
                                        _ = v29
                                    }
                                    _ = v26
                                }
                                var v30 uint16 = 369
                                var v31 bool = false
                                _ = v30
                                _ = v31
                            }
                            v6 = v6
                        } else {
                            v32 := v2[(v1 + v1)]
                            fmt.Println(v6, fi21)
                            if true {
                                fmt.Println(fi21)
                                if len(v15) > 0 {
                                    v15[0] = (fi22 % (((v32 ^ fi22)) | 1))
                                }
                                fmt.Println(v32, v12, v12)
                            }
                            _ = v32
                        }
                        fmt.Println(v1)
                        _ = fi22
                    }
                case -1:
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                }
                var v33 uint16 = 2
                var v34 int8 = int8(v5)
                _ = fi21
                _ = v33
                _ = v34
            }
        }
        _ = len(v15)
        _ = v15
    } else {
        if true {
            var v35 int64 = (v6 + -9223372036854775808)
            v12 = -2.25
            _ = v35
        }
        v36 := GenS5{F0: 150, F1: 127, F2: 2147483647}
        _ = v36
    }
    v37 := []int{303, -647, 32767, 32768}
    delete(v2, "hello")
    var v38 int8 = 127
    for fi39 := 0; fi39 < 17; fi39++ {
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
        _ = fi39
    }
    _ = len(v3)
    if (false || true) {
        switch ("" + (("foo bar" + "\n\t") + (v1 + ("foo bar" + "foo bar")))) {
        case "foo bar":
            v1 = (v1 + (v1 + "hello"))
        default:
            if true {
                var v40 int8 = (v38 - (v38 | (v38 << (((v5 + (^(v5)))) & 63))))
                _ = v40
            }
        }
        for fi41 := 0; fi41 < 19; fi41++ {
            if false {
                if false {
                    v2[("\n\t" + (("" + "\n\t") + v1))] = fi41
                    v12 = v12
                    fmt.Println(fi41)
                } else {
                    fmt.Println(fi41, v5)
                }
                var v42 int64 = v6
                _ = v42
            } else {
                v43 := GenS6{F0: v5, F1: fi41}
                fmt.Println(v1, fi41)
                _ = v43
            }
            v1 = (("hello" + ((v1 + (v1 + ((v1 + ("\n\t" + "hello")) + v1))) + "foo bar")) + "")
            _ = fi41
        }
    }
    v44 := GenS4{F0: 482, F1: 1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
