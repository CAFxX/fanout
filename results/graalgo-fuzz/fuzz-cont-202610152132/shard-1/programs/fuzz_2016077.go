package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
func genFn0() int16 {
    if (!true) {
        for fi6 := 0; fi6 < 18; fi6++ {
            if false {
                fmt.Println(fi6)
            } else {
                v7 := GenS1{F0: -216, F1: 1, F2: 145}
                var v8 float64 = 1.5
                var v9 int64 = 42
                _ = v7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        if false {
            if true {
                v10 := GenS2{F0: 0}
                v11 := []int{336}
                var v12 uint64 = 0
                _ = v10
                _ = v11
                _ = v12
            } else {
                var v13 uint64 = 4294967295
                v14 := []int{-259, 32768}
                _ = v13
                _ = v14
            }
            if true {
                var v15 uint32 = 42
                _ = v15
            }
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                var v16 uint = 2
                _ = v16
            case "":
                var v17 int32 = -759
                _ = v17
            case "a":
                switch -944 {
                case 128:
                    for fi18 := 0; fi18 < 2; fi18++ {
                        if false {
                            var v19 uint8 = 255
                            v19 = v19
                            _ = v19
                        }
                        v20 := GenS3{F0: "hello"}
                        _ = fi18
                        _ = v20
                    }
                case 500:
                    var v21 float64 = 1e10
                    _ = v21
                default:
                    var v22 int64 = 9223372036854775807
                    _ = v22
                }
            default:
                switch (("foo bar" + "") + "foo bar") {
                case "foo bar":
                    var v23 uint = 740
                    _ = v23
                case "\n\t":
                    v24 := []int{}
                    _ = v24
                case "a":
                    var v25 int8 = 20
                    _ = v25
                default:
                    var v26 uint16 = 0
                    _ = v26
                }
            }
        } else {
            var v27 uint = 0
            _ = v27
        }
    }
    var v28 int64 = 0
    v28 = (^(v28))
    v28 = 42
    _ = v28
    return int16(v28)
}

func main() {
    var v1 int64 = 602
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    v4 := []string{"hello"}
    var v5 int32 = 0
    var v29 int16 = 431
    var v30 uint64 = 58
    for fi31 := 0; fi31 < 1; fi31++ {
        if false {
            var v32 float64 = 1.5
            _ = v32
        }
        _ = fi31
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        switch fi33 {
        case 4294967295:
            var v34 float64 = 0.1
            _ = v34
        case 42:
            v1 = (v1 >> ((4294967295) & 63))
        }
        if (((1e10 * -2.25) >= 0.1) && (594 < 945)) {
            if ((v29 + 1) == v29) {
                v5 = v5
                v4 = append(v4, "")
                fmt.Println(v29)
            }
        }
        _ = fi33
    }
    fmt.Println(v5, v30)
    var v35 int = 1
    v36 := GenS3{F0: ("hello" + ("" + "hello"))}
    var v37 uint = 1
    switch v35 {
    case -983:
        fmt.Println(v5, v37)
    case 1:
        var v38 int32 = (v5 + (v5 * v5))
        _ = v38
    }
    if (false && true) {
        fmt.Println(v1)
        var v39 uint16 = 42
        _ = v39
    } else {
        v4 = append(v4, "hello")
        fmt.Println(v1, v30, v1)
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v40 int = (v35 + 128)
    switch (v35 / ((-1) | 1)) {
    case 0:
        if (((v1 & -250) > (v1 / (((v1 * v1)) | 1))) && false) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v41 uint64 = 797
            _ = v41
        }
    case -2147483648:
        fmt.Println(v40, v40, v35)
    case -9223372036854775808:
        var v42 string = "hello"
        _ = v42
    default:
        v43 := GenS1{F0: (v29 / (((v29 & v29)) | 1)), F1: 1, F2: (v30 / (((^(v30))) | 1))}
        _ = v43
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
