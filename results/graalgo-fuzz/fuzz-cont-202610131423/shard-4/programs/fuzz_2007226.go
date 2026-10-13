package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int32; F1 string; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint16; F2 int64 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 16; fi5++ {
        if (fi5 < fi5) {
            var v6 string = ("hello" + ((("" + ("hello" + "")) + "\n\t") + "foo bar"))
            fmt.Println(fi5)
            switch fi5 {
            case 42:
                fmt.Println(fi5)
            default:
                v7 := []int{}
                _ = v7
            }
            _ = v6
        }
        _ = fi5
    }
    if true {
        if (-981 <= 323) {
            if true {
                var v8 float32 = 0.1
                if (!(true && (!true))) {
                    v9 := []int{4294967295, 42, 32768}
                    v10 := GenS1{F0: 1}
                    fmt.Println(v8)
                    _ = v10
                    _ = v9
                }
                _ = v8
            }
            v11 := []int{-1}
            _ = v11
        }
    } else {
        v12 := []int{32768}
        _ = v12
    }
    return -461
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]string)
    var v3 bool = (true && (0 == -128))
    var v4 uint64 = 2
    if (!true) {
        fmt.Println(v3, v1)
        for fi13 := 0; fi13 < 17; fi13++ {
            _ = len(v2)
            if (127 > fi13) {
                v3 = v3
                if ((v1 % (((v1 | (v1 >> (((v1 ^ (^(v1)))) & 63)))) | 1)) != v1) {
                    var v14 uint = (v1 + (^(v1)))
                    _ = v14
                }
            } else {
                v15 := GenS2{F0: v4, F1: 0}
                if (240 != (v4 ^ (v4 & (v4 ^ (^(v4)))))) {
                    v1 -= v1
                    v1 = 18446744073709551615
                }
                var v16 uint32 = 1
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        var v17 string = "foo bar"
        _ = v17
    } else {
        var v18 bool = (0.1 == 1.5)
        _ = v18
    }
    var v19 int32 = -2147483648
    v20 := []int{4294967295}
    _ = len(v2)
    var v21 uint32 = 4294967295
    switch 0 {
    case 634:
        if (!v3) {
            fmt.Println(v21)
            var v22 uint = (v1 - 65535)
            var v23 uint = 851
            _ = v22
            _ = v23
        } else {
            if len(v20) > 0 {
                v20[0] = 9223372036854775807
            }
            fmt.Println(v3, v3, v1)
        }
    case 42:
        for fi24 := 0; fi24 < 16; fi24++ {
            switch fi24 {
            case 127:
                var v25 uint = v1
                _ = v25
            case 9223372036854775807:
                v4 = (v4 & v4)
            case 32768:
                var v26 uint8 = 1
                _ = v26
            default:
                v27 := GenS3{F0: -1, F1: ("a" + ("" + ("foo bar" + ""))), F2: 255}
                _ = v27
            }
            v20 = append(v20, (fi24 - (-(fi24))))
            _ = fi24
        }
    case 568:
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 uint64 = v4
            for fi30 := 0; fi30 < 11; fi30++ {
                switch (((("foo bar" + (((("hello" + "hello") + ("foo bar" + "")) + "foo bar") + "\n\t")) + "") + ("foo bar" + ("hello" + ("\n\t" + "a")))) + "a") {
                case "foo bar":
                    var v31 string = "foo bar"
                    _ = v31
                case "":
                    fmt.Println(v19, fi30, v21)
                case "\n\t":
                    var v32 int32 = 525
                    _ = v32
                default:
                    v2[""] = ("" + ("hello" + "a"))
                }
                switch (("a" + "hello") + ("a" + "a")) {
                case "hello":
                    var v33 int32 = -2147483648
                    _ = v33
                case "foo bar":
                    var v34 int = -904
                    _ = v34
                case "\n\t":
                    v35 := GenS4{F0: -457}
                    _ = v35
                default:
                    if len(v20) > 0 {
                        v20[0] = fi30
                    }
                }
                v3 = v3
                _ = fi30
            }
            if v3 {
                fmt.Println(v21, v21, v4)
            }
            _ = fi28
            _ = v29
        }
    }
    var v36 string = (("a" + ("\n\t" + "foo bar")) + ("" + "\n\t"))
    if (v3 || v3) {
        fmt.Println(v21, v4)
        if false {
            fmt.Println(v36)
            v3 = ((!(uint8(v1) < 128)) && (!false))
            if true {
                for fi37 := 0; fi37 < 2; fi37++ {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        v20 = append(v20, (^(fi37)))
                        _ = fi38
                    }
                    fmt.Println(v1, v21)
                    _ = fi37
                }
            }
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    }
    v39 := GenS5{F0: true, F1: 507, F2: 498}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v39
    _ = v4
}
