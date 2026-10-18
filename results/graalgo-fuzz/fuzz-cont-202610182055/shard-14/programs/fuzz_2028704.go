package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 string; F1 int }

func main() {
    v1 := []string{"\n\t", "\n\t"}
    var v2 bool = ((-2147483648 > -1) && true)
    var v3 int64 = 725
    var v4 int32 = int32(v3)
    v3 += -1
    for fi5 := 0; fi5 < 20; fi5++ {
        if v2 {
            var v6 uint8 = 0
            v4 += v4
            var v7 int = fi5
            _ = v6
            _ = v7
        }
        if true {
            v8 := GenS1{F0: 1000}
            v4 += (v4 ^ -1)
            var v9 uint8 = 198
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    var v10 string = "a"
    if false {
        for fi11 := 0; fi11 < 9; fi11++ {
            v2 = v2
            _ = len(v1)
            switch 9223372036854775807 {
            case 32768:
                if v2 {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        switch v10 {
                        case "\n\t":
                            v13 := GenS1{F0: 1}
                            _ = v13
                        case "a":
                            v1 = append(v1, ("a" + v10))
                        case "":
                            v3 -= (v3 % ((-891) | 1))
                        }
                        v14 := GenS2{F0: 1, F1: v4}
                        if (!(v10 > v10)) {
                            var v15 int16 = -1
                            v2 = v2
                            var v16 bool = true
                            _ = v15
                            _ = v16
                        }
                        _ = fi12
                        _ = v14
                    }
                    v17 := GenS3{F0: "hello", F1: fi11}
                    var v18 uint32 = 65535
                    _ = v17
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        switch ((v10 + v10) + v10) {
                        case "a":
                            var v20 float32 = 3.14159
                            _ = v20
                        case "hello":
                            var v21 uint32 = 2147483648
                            _ = v21
                        case "\n\t":
                            v22 := GenS2{F0: 18446744073709551615, F1: v4}
                            _ = v22
                        }
                        _ = fi19
                    }
                }
            case -1:
                var v23 uint64 = 592
                _ = v23
            case 42:
                fmt.Println(v10, v10)
            }
            _ = fi11
        }
    } else {
        var v24 uint16 = 42
        v25 := GenS3{F0: v10, F1: 912}
        if v2 {
            v24 = (v24 ^ (v24 / ((v24) | 1)))
            var v26 int = -9223372036854775808
            if v2 {
                if (!true) {
                    var v27 uint8 = uint8(v24)
                    _ = v27
                } else {
                    switch v10 {
                    case "\n\t":
                        v28 := GenS2{F0: 0, F1: 132}
                        _ = v28
                    case "":
                        v26 += (-(v26))
                    default:
                        var v29 bool = (float64(float32(((-2.25 * 1.5) + (((0.1 * 1.5) * 3.14159) - -2.25)))) <= 1.5)
                        _ = v29
                    }
                    v10 = "hello"
                }
                v30 := GenS3{F0: ("hello" + ("a" + "hello")), F1: 128}
                _ = v30
            } else {
                var v31 uint32 = 2147483648
                switch "\n\t" {
                case "foo bar":
                    var v32 int16 = -183
                    _ = v32
                case "hello":
                    v33 := GenS1{F0: 1000}
                    _ = v33
                case "\n\t":
                    var v34 uint = 42
                    _ = v34
                default:
                    v2 = v2
                }
                switch v10 {
                case "foo bar":
                    v4 = (v4 & (v4 | (^(v4))))
                case "":
                    v2 = v2
                case "hello":
                    fmt.Println(v10, v2, v24)
                }
                _ = v31
            }
            _ = v26
        }
        _ = v24
        _ = v25
    }
    var v35 float64 = (((3.14159 * float64(3.14159)) / 1e10) / 1.5)
    v35 -= v35
    switch -204 {
    case 42:
        v36 := GenS1{F0: 32767}
        _ = v36
    }
    var v37 int = 4294967295
    fmt.Println(v10)
    v37 = (v37 & v37)
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 int = -733
        _ = fi38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
}
