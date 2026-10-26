package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint8; F2 int }
func genFn0(p0 int32, p1 uint) uint8 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: -1}
        v9 := GenS1{F0: 42}
        var v10 string = "a"
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{int(p0), -900, 42, 509}
    v12 := GenS2{F0: 2147483647, F1: (p0 % ((p0) | 1)), F2: 161}
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint = p1
        var v15 uint = p1
        _ = fi13
        _ = v14
        _ = v15
    }
    if len(v11) > 0 {
        v11[0] = -480
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 128
}

func main() {
    var v1 int16 = -1
    var v2 int8 = -119
    var v3 int = 2147483647
    v4 := make(map[uint32]uint32)
    var v5 int32 = -2147483648
    var v6 int32 = v5
    switch (("foo bar" + "hello") + ("\n\t" + ((("" + (("\n\t" + "") + "")) + "\n\t") + (((("foo bar" + "\n\t") + "") + "a") + ((((("\n\t" + "foo bar") + ((((("\n\t" + "hello") + (("" + (("hello" + ((("" + "a") + "foo bar") + ("foo bar" + ((("foo bar" + "foo bar") + "\n\t") + "\n\t")))) + "foo bar")) + "a")) + ("a" + "a")) + "a") + "\n\t")) + (((("" + "hello") + "") + "\n\t") + "a")) + ("" + "")) + "\n\t"))))) {
    case "\n\t":
        var v16 int16 = (v1 >> ((65535) & 63))
        _ = v16
    default:
        v6 -= (v5 + v5)
    }
    var v17 uint16 = 149
    v6 = v6
    var v18 uint16 = v17
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 string = ("foo bar" + "\n\t")
            _ = fi20
            _ = v21
        }
        switch -2147483648 {
        case 1:
            fmt.Println(v3, fi19)
        }
        _ = fi19
    }
    v18 = 255
    switch "" {
    case "a":
        var v22 uint8 = 31
        _ = v22
    case "foo bar":
        v3 = (v3 * (v3 | (v3 / ((v3) | 1))))
    }
    v2 += -128
    var v23 string = ""
    if (v5 == v6) {
        v24 := GenS1{F0: 32767}
        switch 32768 {
        case 0:
            for fi25 := 0; fi25 < 13; fi25++ {
                switch (v23 + ("\n\t" + "hello")) {
                case "hello":
                    var v26 uint64 = 397
                    _ = v26
                case "":
                    var v27 bool = (!true)
                    _ = v27
                default:
                    v28 := GenS2{F0: v6, F1: -894, F2: 0}
                    _ = v28
                }
                v29 := GenS3{F0: -9223372036854775808}
                _ = fi25
                _ = v29
            }
        }
        v3 = v3
        _ = v24
    } else {
        if (662 != 2) {
            v3 = v3
            for fi30 := 0; fi30 < 7; fi30++ {
                if true {
                    fmt.Println(v6, v17, fi30)
                    v18 += (v18 - (v18 / ((32767) | 1)))
                    v5 = (v6 + (^(v5)))
                } else {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        switch v23 {
                        case "":
                            v5 = v6
                            var v32 int16 = 985
                            switch ("a" + ("a" + "hello")) {
                            case "a":
                                if (!(!true)) {
                                    var v33 int32 = -775
                                    _ = v33
                                }
                                var v34 uint = 788
                                _ = v34
                            }
                            _ = v32
                        default:
                            v35 := GenS4{F0: v18}
                            _ = v35
                        }
                        v18 = v17
                        v36 := v4[2147483648]
                        _ = fi31
                        _ = v36
                    }
                }
                _ = fi30
            }
        }
        v37 := GenS2{F0: -856, F1: v5, F2: 53}
        fmt.Println(v5, v5, v5)
        _ = v37
    }
    v5 = -763
    for fi38 := 0; fi38 < 4; fi38++ {
        fmt.Println(v6)
        switch (-(v3)) {
        case 0:
            var v39 int16 = int16(v6)
            _ = v39
        case 2147483647:
            v40 := v4[0]
            _ = v40
        }
        switch fi38 {
        case 32767:
            for fi41 := 0; fi41 < 13; fi41++ {
                var v42 int8 = (v2 - 1)
                _ = fi41
                _ = v42
            }
        case 32768:
            if (1 == v5) {
                _ = len(v4)
                fmt.Println(v23, v3, v3)
            }
        case 0:
            if (!true) {
                v43 := []int{v3, -2147483648, int(v2), (fi38 * fi38)}
                _ = v43
            } else {
                v44 := []int{2147483647, 4294967295, v3}
                _ = v44
            }
        default:
            v45 := []int{-1, (fi38 % ((v3) | 1)), (v3 & 4294967295), (fi38 % ((42) | 1))}
            _ = v45
        }
        _ = fi38
    }
    v46 := v4[42]
    v47 := GenS5{F0: 340, F1: 164, F2: 915}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
