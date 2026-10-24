package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 0.0
    var v3 int16 = 966
    v4 := make(map[string]string)
    var v5 uint32 = 4294967295
    var v6 uint16 = 255
    var v7 float32 = 3.14159
    v8 := []int{-9223372036854775808}
    v4[(("hello" + (("" + "\n\t") + "")) + "hello")] = (("foo bar" + "\n\t") + "hello")
    var v9 uint16 = v6
    var v10 float32 = v7
    if true {
        fmt.Println(v9)
        v11 := v4[""]
        _ = v11
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            v5 += 2147483647
            _ = fi12
        }
    }
    v13 := GenS1{F0: 91, F1: 411}
    v14 := GenS1{F0: 127, F1: -1}
    if true {
        for fi15 := 0; fi15 < 1; fi15++ {
            v10 -= v10
            _ = fi15
        }
        switch 2147483647 {
        case 127:
            switch ("\n\t" + ("foo bar" + "\n\t")) {
            case "a":
                var v16 uint16 = (v6 >> ((4294967295) & 63))
                _ = v16
            case "":
                switch (("" + "a") + "foo bar") {
                case "":
                    v9 += v9
                default:
                    for fi17 := 0; fi17 < 15; fi17++ {
                        v3 -= 105
                        for fi18 := 0; fi18 < 2; fi18++ {
                            v19 := GenS1{F0: 166, F1: 42}
                            if true {
                                v20 := GenS1{F0: 2, F1: -9223372036854775808}
                                fmt.Println(fi18)
                                var v21 uint8 = 22
                                _ = v20
                                _ = v21
                            } else {
                                if false {
                                    if true {
                                        delete(v1, ("a" + ("a" + ("" + "foo bar"))))
                                        switch "\n\t" {
                                        case "\n\t":
                                            fmt.Println(v6, v5, v5)
                                            switch (((("" + "") + "a") + "hello") + "\n\t") {
                                            case "foo bar":
                                                v22 := GenS1{F0: 174, F1: -1}
                                                _ = v22
                                            case "hello":
                                                var v23 bool = (!(127 != 42))
                                                _ = v23
                                            case "":
                                                _ = len(v1)
                                            default:
                                                v24 := GenS1{F0: 206, F1: -829}
                                                _ = v24
                                            }
                                        case "foo bar":
                                            switch "a" {
                                            case "foo bar":
                                                var v25 uint16 = v6
                                                _ = v25
                                            case "a":
                                                var v26 uint = 65535
                                                _ = v26
                                            default:
                                                switch "a" {
                                                case "foo bar":
                                                    switch ("hello" + "\n\t") {
                                                    case "a":
                                                        v27 := GenS2{F0: 531, F1: 0}
                                                        _ = v27
                                                    case "hello":
                                                        v1[""] = 65535
                                                    }
                                                case "a":
                                                    v8 = append(v8, -340)
                                                }
                                            }
                                        default:
                                            v5 -= v5
                                        }
                                    }
                                    v3 -= v3
                                }
                                if true {
                                    if (!true) {
                                        v3 = (v3 >> ((1) & 63))
                                        var v28 uint8 = 255
                                        if (!((!false) || (288 == 255))) {
                                            v6 = (^(v9))
                                        }
                                        _ = v28
                                    } else {
                                        v10 -= 3.14159
                                    }
                                    fmt.Println(fi17)
                                }
                            }
                            var v29 uint32 = v5
                            v7 = ((v7 * (3.14159 + (3.14159 + ((-2.25 * float32(3.14159)) + (0.0 + -2.25))))) - (v7 - -2.25))
                            _ = fi18
                            _ = v19
                            _ = v29
                        }
                        var v30 float32 = (v10 + ((v10 / float32(v2)) * v7))
                        _ = fi17
                        _ = v30
                    }
                }
            case "hello":
                if true {
                    v4["hello"] = "a"
                } else {
                    var v31 float32 = ((v10 / (v10 * ((v10 - (float32((v2 - v2)) / v10)) - ((0.0 + (0.0 * -2.25)) * v10)))) + v10)
                    _ = v31
                }
            default:
                fmt.Println(v3, v7, v6)
            }
        case 1:
            v32 := GenS3{F0: v6, F1: 2, F2: (v3 >> ((18446744073709551615) & 63))}
            _ = v32
        }
    } else {
        v33 := v4["foo bar"]
        _ = v33
    }
    v6 -= (v6 | 42)
    for fi34 := 0; fi34 < 13; fi34++ {
        switch "" {
        case "foo bar":
            v10 = 0.0
        default:
            var v35 uint8 = 107
            _ = v35
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi34
    }
    var v36 int8 = 42
    var v37 int16 = v3
    for fi38 := 0; fi38 < 13; fi38++ {
        fmt.Println(fi38, v36, v5)
        fmt.Println(v10, v7)
        _ = len(v8)
        _ = fi38
    }
    var v39 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
