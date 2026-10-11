package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int32; F1 int64; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 int16 = 32767
    v4 := make(map[uint32]bool)
    v5 := []string{(("" + "hello") + (("a" + ("a" + "foo bar")) + (("\n\t" + ("a" + "foo bar")) + ("foo bar" + "")))), "a", "hello", ("" + "")}
    v3 -= (v3 >> ((0) & 63))
    v3 += v3
    v3 = (v3 - v3)
    for fi6 := 0; fi6 < 17; fi6++ {
        v3 = (v3 >> ((255) & 63))
        fmt.Println(fi6, v3)
        _ = fi6
    }
    switch 32768 {
    case 1:
        var v7 int64 = -801
        _ = v7
    case -2147483648:
        v3 = -1
    }
    var v8 bool = true
    for fi9 := 0; fi9 < 10; fi9++ {
        v3 -= -32768
        v8 = true
        var v10 uint = 0
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    var v12 bool = ((42 <= 0) && (255 == 1))
    switch 602 {
    case 9223372036854775807:
        switch -658 {
        case 1:
            fmt.Println(v8, v3)
        case 21:
            for fi13 := 0; fi13 < 14; fi13++ {
                if (!(1 > 620)) {
                    v4[2147483647] = ((0.0 >= (1.5 - -2.25)) && v12)
                    fmt.Println(v8)
                    v14 := GenS1{F0: 39}
                    _ = v14
                } else {
                    var v15 float64 = 3.14159
                    v16 := GenS1{F0: 125}
                    _ = v15
                    _ = v16
                }
                fmt.Println(v12)
                v3 -= v3
                _ = fi13
            }
        case 642:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            v17 := v2[4294967295]
            _ = v17
        }
    default:
        v3 = v3
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        switch "" {
        case "hello":
            for fi19 := 0; fi19 < 3; fi19++ {
                if (("a" + "") > "hello") {
                    _ = len(v5)
                }
                if true {
                    if false {
                        fmt.Println(v8, v8)
                        v20 := GenS2{F0: (^(v3)), F1: ("hello" + ("hello" + (("" + ("a" + "")) + ("hello" + (("" + ((("" + "a") + "foo bar") + "a")) + "foo bar")))))}
                        _ = len(v5)
                        _ = v20
                    } else {
                        switch (-(fi19)) {
                        case -670:
                            v8 = (!((v3 << ((4294967295) & 63)) > (v3 / ((v3) | 1))))
                        case -2147483648:
                            v8 = ("\n\t" <= "")
                        default:
                            var v21 uint8 = 2
                            _ = v21
                        }
                        var v22 uint = 1
                        _ = v22
                    }
                    v23 := v4[2147483648]
                    if (v3 > 0) {
                        var v24 uint16 = 255
                        if v23 {
                            var v25 int8 = 1
                            _ = v25
                        } else {
                            v8 = false
                        }
                        _ = v24
                    } else {
                        fmt.Println(fi18)
                        v8 = v8
                        var v26 int = fi18
                        _ = v26
                    }
                    v27 := GenS1{F0: 0}
                    _ = v23
                    _ = v27
                } else {
                    var v28 int8 = -54
                    var v29 uint8 = 1
                    v5 = append(v5, (("hello" + ("a" + "a")) + "a"))
                    _ = v28
                    _ = v29
                }
                _ = fi19
            }
        case "":
            switch (("" + "a") + "foo bar") {
            case "":
                for fi30 := 0; fi30 < 9; fi30++ {
                    fmt.Println(v8, fi30, fi18)
                    _ = fi30
                }
            default:
                for fi31 := 0; fi31 < 3; fi31++ {
                    v32 := GenS2{F0: v3, F1: "\n\t"}
                    if false {
                        v8 = false
                        var v33 int = 2147483647
                        if v8 {
                            v12 = false
                        }
                        _ = v33
                    }
                    fmt.Println(fi18, v8)
                    _ = fi31
                    _ = v32
                }
            }
        }
        var v34 string = "a"
        _ = fi18
        _ = v34
    }
    v4[374] = (!((v3 % ((v3) | 1)) > 42))
    if v12 {
        for fi35 := 0; fi35 < 12; fi35++ {
            for fi36 := 0; fi36 < 11; fi36++ {
                v12 = v12
                _ = fi36
            }
            switch (fi35 * fi35) {
            case 0:
                var v37 uint16 = 255
                _ = v37
            default:
                v38 := GenS3{F0: 2147483647, F1: 216, F2: 63}
                _ = v38
            }
            var v39 float32 = float32(1.5)
            _ = fi35
            _ = v39
        }
        _ = len(v5)
    } else {
        v12 = (2147483647 != 1)
    }
    v12 = (3.14159 == 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
