package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }

func main() {
    v1 := []string{"a"}
    var v2 int32 = 207
    v3 := []int{32768}
    var v4 float64 = (0.0 + float64((0.0 * float32(float64(1e10)))))
    var v5 int16 = 32767
    if true {
        v4 = (float64(1.5) * ((v4 - v4) / v4))
    } else {
        switch 42 {
        case 42:
            v2 = v2
        case -2147483648:
            if true {
                v5 = (v5 | v5)
                var v6 bool = (0 != 4294967295)
                var v7 float64 = v4
                _ = v6
                _ = v7
            }
        }
    }
    var v8 uint16 = 255
    if false {
        v2 = v2
        var v9 int64 = -1
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int16 = v5
            switch (("hello" + "") + ("" + "a")) {
            case "a":
                v12 = (v5 << ((2) & 63))
            case "foo bar":
                v13 := GenS1{F0: (0.1 >= ((v4 - float64(0.1)) * 0.0)), F1: 18}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v14 float32 = (0.0 + float32(float64((1.5 - -2.25))))
                _ = v13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        switch fi10 {
        case 9223372036854775807:
            var v15 uint64 = 708
            _ = v15
        default:
            fmt.Println(fi10, v4)
        }
        _ = fi10
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v5)
        switch fi16 {
        case -9223372036854775808:
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        case -1:
            v17 := GenS2{F0: (!true)}
            _ = v17
        case -2147483648:
            switch "foo bar" {
            case "":
                fmt.Println(v8)
            default:
                v8 = v8
            }
        default:
            switch ((("" + "a") + "foo bar") + ("foo bar" + ("hello" + ("\n\t" + "foo bar")))) {
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "hello":
                switch (fi16 >> ((255) & 63)) {
                case 366:
                    var v18 uint = 4294967295
                    _ = v18
                case 0:
                    fmt.Println(v5, v5)
                case -9223372036854775808:
                    for fi19 := 0; fi19 < 1; fi19++ {
                        var v20 int = (fi19 - fi19)
                        if false {
                            if (true || (2147483648 != 1)) {
                                v21 := GenS3{F0: (!(!true))}
                                _ = v21
                            } else {
                                var v22 int = (v20 - -1)
                                var v23 bool = ((!(true && true)) && false)
                                fmt.Println(v2, v2, v23)
                                _ = v22
                                _ = v23
                            }
                        } else {
                            switch fi16 {
                            case 9223372036854775807:
                                var v24 uint = 4294967295
                                _ = v24
                            case -827:
                                fmt.Println(fi19, v8, v2)
                            }
                            var v25 uint = 159
                            var v26 int16 = -929
                            _ = v25
                            _ = v26
                        }
                        v4 = (v4 - v4)
                        _ = fi19
                        _ = v20
                    }
                default:
                    var v27 int8 = 100
                    _ = v27
                }
            default:
                var v28 int64 = 39
                _ = v28
            }
        }
        _ = fi16
    }
    if false {
        for fi29 := 0; fi29 < 18; fi29++ {
            v8 = v8
            _ = fi29
        }
    }
    v8 += v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
