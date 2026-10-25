package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 65535
    v3 := []int{4294967295}
    var v4 float32 = 0.1
    v5 := make(map[uint32]string)
    var v6 int16 = 32767
    fmt.Println(v6)
    var v7 string = "hello"
    v8 := GenS1{F0: 0, F1: 32767, F2: -1}
    v6 -= v6
    v9 := GenS2{F0: 18446744073709551615, F1: 127, F2: v7}
    if (((!true) && ((v2 > 0) || (290 == 4294967295))) && (4294967295 < 42)) {
        v10 := GenS2{F0: 42, F1: 255, F2: (((("foo bar" + v7) + "") + "") + "foo bar")}
        _ = v10
    } else {
        fmt.Println(v2, v7, v7)
        if false {
            v6 -= v6
            for fi11 := 0; fi11 < 17; fi11++ {
                if true {
                    if (0.0 != ((1.5 * (0.0 - 1e10)) - (-2.25 * 0.1))) {
                        for fi12 := 0; fi12 < 18; fi12++ {
                            _ = len(v1)
                            _ = fi12
                        }
                        v3 = append(v3, 32768)
                    } else {
                        for fi13 := 0; fi13 < 18; fi13++ {
                            v14 := GenS2{F0: 0, F1: 0, F2: v7}
                            v2 -= v2
                            _ = fi13
                            _ = v14
                        }
                        var v15 uint64 = uint64(v2)
                        if true {
                            for fi16 := 0; fi16 < 17; fi16++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v17 uint64 = (v15 >> ((972) & 63))
                                _ = fi16
                                _ = v17
                            }
                            var v18 string = v7
                            _ = len(v3)
                            _ = v18
                        }
                        _ = v15
                    }
                } else {
                    v2 = v2
                    var v19 string = (("hello" + v7) + "foo bar")
                    _ = v19
                }
                _ = fi11
            }
            switch 2147483647 {
            case 567:
                _ = len(v3)
            case 32768:
                for fi20 := 0; fi20 < 4; fi20++ {
                    v2 += 84
                    _ = fi20
                }
            default:
                switch v7 {
                case "hello":
                    v7 = v7
                case "foo bar":
                    var v21 uint8 = 42
                    _ = v21
                case "\n\t":
                    if false {
                        v6 = v6
                        v22 := v5[v2]
                        _ = len(v3)
                        _ = v22
                    }
                }
            }
        }
        v23 := v1[-2147483648]
        _ = v23
    }
    switch "a" {
    case "a":
        v6 += -1
    case "hello":
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := GenS2{F0: 633, F1: 255, F2: "a"}
            _ = fi24
            _ = v25
        }
    }
    switch v7 {
    case "\n\t":
        v7 = "foo bar"
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        if true {
            v4 = v4
        } else {
            for fi27 := 0; fi27 < 2; fi27++ {
                var v28 string = ("hello" + ((((v7 + "\n\t") + (v7 + v7)) + "a") + v7))
                var v29 uint16 = 2
                v30 := GenS2{F0: 65535, F1: 38, F2: "a"}
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
        switch "\n\t" {
        case "hello":
            _ = len(v3)
        case "a":
            if ((v2 - (v2 % (((v2 | v2)) | 1))) == 2147483647) {
                var v31 uint32 = 4294967295
                fmt.Println(v7, v2)
                _ = v31
            }
        case "":
            v32 := v1[fi26]
            _ = v32
        default:
            if len(v3) > 0 {
                v3[0] = (fi26 >> ((18446744073709551615) & 63))
            }
        }
        if (249 == -1) {
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 int8 = 0
                v4 += float32(float64((-2.25 + (v4 * float32(-2.25)))))
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 3; fi35++ {
                v7 = "hello"
                var v36 bool = true
                fmt.Println(fi35, v7)
                v7 = v7
                _ = fi35
                _ = v36
            }
            var v37 bool = true
            _ = v37
        }
        _ = fi26
    }
    v38 := GenS3{F0: -2147483648, F1: v6}
    v39 := GenS4{F0: 2}
    var v40 bool = false
    var v41 int8 = 0
    var v42 uint64 = 42
    fmt.Println(v41, v7, v42)
    var v43 string = ("" + (v7 + v7))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
