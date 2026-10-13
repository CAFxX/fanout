package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"hello", "foo bar", ("hello" + ("foo bar" + "\n\t")), ("a" + "a")}
    var v3 string = "\n\t"
    v4 := make(map[int]int)
    if false {
        v4[32768] = 128
        var v5 uint = 65535
        _ = v5
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        switch (("\n\t" + v3) + "") {
        case "\n\t":
            if true {
                var v7 string = ((("" + (v3 + v3)) + v3) + "foo bar")
                v8 := GenS1{F0: 615, F1: 255}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v7
                _ = v8
            } else {
                switch fi6 {
                case 1:
                    v3 = v3
                case 32767:
                    var v9 uint32 = 1
                    _ = v9
                case -1:
                    fmt.Println(fi6, v3)
                default:
                    var v10 float32 = 3.14159
                    _ = v10
                }
            }
        case "foo bar":
            switch v3 {
            case "a":
                switch fi6 {
                case -557:
                    if true {
                        v11 := GenS2{F0: 0}
                        _ = v11
                    }
                default:
                    var v12 uint8 = 127
                    _ = v12
                }
            case "hello":
                var v13 int32 = 778
                _ = v13
            case "\n\t":
                v3 = ""
            }
        }
        if (!(!(true || true))) {
            v14 := GenS2{F0: 508}
            if (!false) {
                v3 = v3
                v4[fi6] = (fi6 ^ fi6)
                fmt.Println(fi6, v3)
            } else {
                switch "" {
                case "hello":
                    var v15 bool = false
                    _ = v15
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    v3 = ((("" + (("foo bar" + v3) + "a")) + (("a" + ("" + (("hello" + (v3 + v3)) + (v3 + "hello")))) + v3)) + (v3 + ""))
                    _ = fi16
                }
                var v17 int8 = -128
                _ = v17
            }
            fmt.Println(fi6, v3)
            _ = v14
        }
        _ = fi6
    }
    var v18 int32 = 420
    v19 := GenS2{F0: v18}
    for fi20 := 0; fi20 < 5; fi20++ {
        if len(v2) > 0 {
            v2[0] = ("foo bar" + v3)
        }
        fmt.Println(v18)
        var v21 uint = 2
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: 247}
    if true {
        for fi23 := 0; fi23 < 14; fi23++ {
            v3 = "\n\t"
            v3 = v3
            var v24 int32 = (^(v18))
            _ = fi23
            _ = v24
        }
        v25 := GenS2{F0: 1}
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 string = "hello"
            fmt.Println(v18, fi26, v3)
            v27 = ((((v3 + (v3 + v27)) + "hello") + v27) + v27)
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        if (133 < 65535) {
            if false {
                if (-853 == 4294967295) {
                    var v28 int8 = 0
                    _ = v28
                }
                v29 := GenS3{F0: 71}
                _ = v29
            }
            var v30 int = -1
            _ = v30
        } else {
            for fi31 := 0; fi31 < 2; fi31++ {
                if ((fi31 + (fi31 & (fi31 | fi31))) > fi31) {
                    v18 -= -2147483648
                    if ((true || false) && false) {
                        v32 := GenS1{F0: 235, F1: 42}
                        v3 = v3
                        _ = v32
                    } else {
                        var v33 uint8 = 42
                        var v34 uint16 = 65535
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v33
                        _ = v34
                    }
                }
                _ = fi31
            }
        }
        var v35 string = ((v3 + v3) + v3)
        fmt.Println(v18, v18, v18)
        _ = v35
    }
    var v36 uint8 = 42
    v18 = (-(v18))
    var v37 int = 32767
    if true {
        v18 = 2147483647
        if (-9223372036854775808 <= -9223372036854775808) {
            fmt.Println(v18, v36)
            var v38 uint = 845
            if true {
                delete(v1, 4294967295)
                var v39 uint32 = 595
                _ = v39
            }
            _ = v38
        }
        var v40 int64 = -1
        _ = v40
    }
    v41 := GenS3{F0: 993}
    v3 = "hello"
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v41
}
