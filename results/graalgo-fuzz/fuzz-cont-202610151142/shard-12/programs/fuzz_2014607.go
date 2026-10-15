package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint8, p2 int32) int8 {
    p0 -= 1
    p0 += 242
    v5 := []int{32768, 9223372036854775807}
    for fi6 := 0; fi6 < 12; fi6++ {
        v5 = append(v5, (fi6 << ((18446744073709551615) & 63)))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (2147483647 >= 204)
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]string)
    fmt.Println(v3)
    for fi7 := 0; fi7 < 17; fi7++ {
        switch fi7 {
        case 32768:
            fmt.Println(v2, v2, fi7)
        case 677:
            var v8 uint = 815
            _ = v8
        }
        if v2 {
            switch (^(fi7)) {
            case 128:
                fmt.Println(fi7, v2, v3)
            case 4294967295:
                fmt.Println(v3)
            case 310:
                switch fi7 {
                case 272:
                    for fi9 := 0; fi9 < 20; fi9++ {
                        v10 := GenS1{F0: 224}
                        _ = fi9
                        _ = v10
                    }
                case 128:
                    if (53 <= 255) {
                        v2 = (!false)
                    }
                default:
                    for fi11 := 0; fi11 < 2; fi11++ {
                        switch (^(fi7)) {
                        case 762:
                            var v12 uint = 680
                            _ = v12
                        default:
                            var v13 uint64 = 84
                            _ = v13
                        }
                        var v14 string = "foo bar"
                        v2 = (((281 != 330) && true) && v2)
                        v15 := GenS1{F0: 255}
                        _ = fi11
                        _ = v14
                        _ = v15
                    }
                }
            }
            var v16 int64 = v3
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := []int{fi17, fi17}
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        v2 = (false && true)
        _ = fi7
    }
    v19 := []int{1}
    var v20 bool = (42 != 26)
    v3 = v3
    if (242 > 2) {
        var v21 string = "foo bar"
        if v20 {
            for fi22 := 0; fi22 < 9; fi22++ {
                for fi23 := 0; fi23 < 6; fi23++ {
                    v2 = (3.14159 > 3.14159)
                    v2 = ((1 <= 1) && v2)
                    if true {
                        v2 = false
                    }
                    _ = fi23
                }
                _ = fi22
            }
            v24 := GenS2{F0: 128}
            _ = v24
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS3{F0: (v21 + ""), F1: 42, F2: 737}
                var v28 int32 = 0
                var v29 int32 = 922
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            if len(v19) > 0 {
                v19[0] = fi25
            }
            for fi30 := 0; fi30 < 4; fi30++ {
                v20 = (v20 || v20)
                v31 := GenS4{F0: 0}
                v2 = ("\n\t" >= "a")
                var v32 uint8 = 93
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi25
        }
        _ = v21
    }
    v19 = append(v19, 2147483647)
    if true {
        for fi33 := 0; fi33 < 3; fi33++ {
            v34 := GenS4{F0: 779}
            v35 := GenS5{F0: 808, F1: "a"}
            _ = fi33
            _ = v34
            _ = v35
        }
    } else {
        v36 := GenS6{F0: (v3 * 1), F1: ("foo bar" + "a")}
        v3 += (-(v3))
        _ = v36
    }
    switch "hello" {
    case "a":
        v37 := GenS7{F0: 96}
        _ = v37
    case "\n\t":
        v1[192] = (37 == -48)
    }
    v20 = v20
    var v38 float32 = 1.5
    switch "a" {
    case "foo bar":
        v39 := GenS5{F0: 4294967295, F1: "a"}
        _ = v39
    }
    v20 = (v3 >= (v3 / ((v3) | 1)))
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    if (!(!v2)) {
        fmt.Println(v2)
    }
    var v40 uint32 = 4294967295
    v20 = (3.14159 <= 1.5)
    switch 128 {
    case 888:
        v41 := GenS1{F0: 108}
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
    _ = v40
}
