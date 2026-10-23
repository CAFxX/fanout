package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 bool }

func main() {
    var v1 uint16 = 65535
    var v2 uint = 4294967295
    var v3 string = "a"
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 6; fi6++ {
        _ = len(v4)
        v3 = v3
        v3 = v3
        _ = fi6
    }
    v7 := GenS1{F0: 58}
    fmt.Println(v2)
    v8 := GenS1{F0: 1}
    var v9 string = ("" + v3)
    fmt.Println(v3, v9, v3)
    if false {
        if true {
            if (!false) {
                switch 32767 {
                case 0:
                    var v10 uint = v2
                    _ = v10
                }
            } else {
                switch 2147483647 {
                case -2147483648:
                    switch "hello" {
                    case "\n\t":
                        for fi11 := 0; fi11 < 8; fi11++ {
                            v3 = ""
                            _ = fi11
                        }
                    case "a":
                        v5[-236] = false
                    default:
                        for fi12 := 0; fi12 < 19; fi12++ {
                            if true {
                                switch v3 {
                                case "":
                                    if true {
                                        if ((v9 + v9) > ("hello" + (v9 + "hello"))) {
                                            fmt.Println(v2, v9, fi12)
                                            fmt.Println(v3)
                                        }
                                        var v13 int32 = -71
                                        _ = v13
                                    }
                                }
                            }
                            _ = fi12
                        }
                    }
                case -554:
                    var v14 float32 = 1e10
                    _ = v14
                case 128:
                    v2 = (v2 | (v2 >> (((^(v2))) & 63)))
                default:
                    v15 := GenS2{F0: -601, F1: 32768}
                    _ = v15
                }
                var v16 float32 = (1e10 / 3.14159)
                _ = v16
            }
            v17 := v5[-2147483648]
            _ = v17
        }
    } else {
        var v18 int16 = 1000
        _ = v18
    }
    if true {
        v5[-64] = false
    } else {
        v9 = v9
    }
    switch (v3 + "a") {
    case "foo bar":
        var v19 uint64 = 127
        _ = v19
    }
    _ = len(v4)
    var v20 bool = (2 <= 128)
    if (v2 < 1) {
        if v20 {
            v3 = "foo bar"
        }
    } else {
        if (v20 || v20) {
            v21 := []int{-655, 4294967295, 2147483647, 9223372036854775807}
            v22 := GenS1{F0: 83}
            for fi23 := 0; fi23 < 8; fi23++ {
                v24 := GenS3{F0: (v9 + ("hello" + (v9 + "foo bar")))}
                v25 := GenS2{F0: -587, F1: (fi23 - -858)}
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v21
            _ = v22
        }
        var v26 float64 = -2.25
        _ = v26
    }
    switch 1 {
    case 9223372036854775807:
        if ((false || (v20 || (v2 >= (v2 / ((v2) | 1))))) || (!true)) {
            v27 := []int{4294967295}
            _ = v27
        } else {
            var v28 uint64 = 1
            var v29 float64 = 1.5
            _ = v28
            _ = v29
        }
    case -920:
        v30 := v4[2147483647]
        _ = v30
    default:
        var v31 int16 = 1000
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 float32 = 3.14159
        for fi34 := 0; fi34 < 20; fi34++ {
            switch fi34 {
            case -335:
                var v35 uint = (v2 >> (((v2 >> ((v2) & 63))) & 63))
                _ = v35
            default:
                v9 = (v9 + "")
            }
            v36 := []int{-1, fi34, fi34, (-(fi34))}
            v37 := GenS2{F0: -32768, F1: 32768}
            _ = fi34
            _ = v36
            _ = v37
        }
        var v38 string = v3
        _ = fi32
        _ = v33
        _ = v38
    }
    if (42 > 222) {
        for fi39 := 0; fi39 < 10; fi39++ {
            switch ("foo bar" + v3) {
            case "hello":
                v40 := GenS4{F0: 1000, F1: v20}
                _ = v40
            case "\n\t":
                fmt.Println(fi39)
            case "foo bar":
                var v41 bool = (249 > 185)
                _ = v41
            default:
                if v20 {
                    v20 = true
                    v42 := []int{fi39, (fi39 << ((42) & 63))}
                    _ = v42
                } else {
                    var v43 int8 = 100
                    var v44 uint64 = 612
                    _ = v43
                    _ = v44
                }
            }
            v45 := GenS2{F0: 1, F1: fi39}
            _ = fi39
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
