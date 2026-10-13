package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -72
    v3 := []string{}
    v4 := []string{}
    v5 := []int{32768, 128}
    var v6 string = ""
    v2 -= (v2 - (v2 - -1))
    _ = len(v1)
    var v7 int = 1
    if (((((42 <= 0) || (0.1 <= (0.0 * ((1.5 * 0.0) + ((((0.0 * ((-2.25 * float64(0.0)) + 3.14159)) - 0.0) * ((float64(((0.0 - 0.0) / (-2.25 / 1.5))) * ((float64(1.5) - float64(float32(3.14159))) + 3.14159)) * 1.5)) * (-2.25 * (3.14159 * (-2.25 + float64(1e10))))))))) && (false || (v7 != (v7 | 128)))) && false) || false) {
        for fi8 := 0; fi8 < 1; fi8++ {
            _ = len(v4)
            switch "hello" {
            case "hello":
                switch ((("hello" + ((("\n\t" + "\n\t") + ("" + (("a" + ("a" + (v6 + (v6 + (v6 + "foo bar"))))) + (v6 + (v6 + v6))))) + "")) + v6) + v6) {
                case "":
                    if ((v2 ^ (^(v2))) <= (v2 & 42)) {
                        for fi9 := 0; fi9 < 18; fi9++ {
                            var v10 float64 = float64(float32((float64(((1.5 / (((1e10 - 0.0) / 1.0) - (-2.25 - (3.14159 - 0.0)))) / (1e10 / 1e10))) / 1.0)))
                            _ = fi9
                            _ = v10
                        }
                        for fi11 := 0; fi11 < 11; fi11++ {
                            fmt.Println(fi8, fi8)
                            _ = fi11
                        }
                        switch "hello" {
                        case "":
                            if false {
                                v1["foo bar"] = ((((v6 + v6) + v6) >= (v6 + v6)) && true)
                                fmt.Println(v2, fi8)
                            }
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = v6
                            }
                        default:
                            v5 = append(v5, v7)
                        }
                    }
                case "\n\t":
                    if (-980 <= 2147483647) {
                        v2 -= (-(v2))
                        if true {
                            v12 := GenS1{F0: fi8, F1: 0}
                            var v13 uint8 = 145
                            _ = v12
                            _ = v13
                        } else {
                            var v14 uint16 = 20
                            var v15 int8 = (-(v2))
                            v1["foo bar"] = true
                            _ = v14
                            _ = v15
                        }
                    }
                default:
                    var v16 uint64 = 1
                    _ = v16
                }
            case "a":
                fmt.Println(v2, v6, v7)
            }
            _ = fi8
        }
        switch ("hello" + "") {
        case "foo bar":
            var v17 float64 = ((-2.25 - (3.14159 * -2.25)) - 1.5)
            _ = v17
        case "hello":
            v6 = v6
        }
    } else {
        v6 = v6
        if (false || true) {
            for fi18 := 0; fi18 < 8; fi18++ {
                v2 += v2
                delete(v1, ("\n\t" + v6))
                _ = fi18
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                _ = len(v1)
                v20 := GenS1{F0: fi19, F1: 1}
                _ = fi19
                _ = v20
            }
            fmt.Println(v2, v6)
        }
        var v21 bool = false
        _ = v21
    }
    v7 = (^(v7))
    var v22 int32 = 0
    switch "" {
    case "":
        fmt.Println(v6, v2, v2)
    }
    v22 = v22
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 13; fi24++ {
            if true {
                var v25 string = (v6 + v6)
                fmt.Println(v2)
                var v26 float64 = 1e10
                _ = v25
                _ = v26
            }
            v4 = append(v4, "\n\t")
            switch (fi24 & v7) {
            case 128:
                var v27 uint16 = 255
                _ = v27
            default:
                switch v6 {
                case "foo bar":
                    if (((v6 + ("" + (("hello" + v6) + "a"))) + ("a" + v6)) >= v6) {
                        switch (fi23 | v7) {
                        case -1:
                            if false {
                                fmt.Println(v6, fi24)
                                switch v6 {
                                case "\n\t":
                                    var v28 uint = 4294967295
                                    _ = v28
                                default:
                                    v29 := GenS1{F0: fi24, F1: 147}
                                    _ = v29
                                }
                            }
                        default:
                            var v30 float64 = 0.1
                            _ = v30
                        }
                        if (0 < 2) {
                            var v31 uint64 = 0
                            _ = v31
                        }
                    } else {
                        fmt.Println(fi23, fi24, fi24)
                    }
                case "hello":
                    var v32 string = v6
                    _ = v32
                default:
                    v6 = (((v6 + v6) + "") + v6)
                }
                v33 := GenS2{F0: 42, F1: false}
                _ = v33
            }
            _ = len(v1)
            v34 := GenS3{F0: 84}
            _ = fi24
            _ = v34
        }
        if false {
            var v35 int64 = int64(v7)
            var v36 uint32 = 2147483648
            _ = v35
            _ = v36
        }
        _ = fi23
    }
    v6 = ""
    v37 := GenS2{F0: 0, F1: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
