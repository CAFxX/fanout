package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int8 = -1
    var v2 uint8 = 182
    v3 := make(map[int]uint32)
    v4 := []string{(("" + ("\n\t" + "")) + ((("\n\t" + ((((("hello" + (("foo bar" + ("" + "")) + "foo bar")) + "") + "hello") + ("" + (("" + "") + "\n\t"))) + "")) + "\n\t") + "a")), (("a" + ("" + "foo bar")) + "a")}
    v5 := make(map[int]string)
    switch -9223372036854775808 {
    case -599:
        var v6 uint32 = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = 42
        switch (-(fi7)) {
        case 351:
            delete(v3, (^(fi7)))
        case -633:
            v9 := GenS1{F0: 578}
            _ = v9
        case 127:
            switch (fi7 / (((fi7 >> ((1) & 63))) | 1)) {
            case -2147483648:
                var v10 int16 = 0
                _ = v10
            default:
                fmt.Println(v8, fi7, v8)
            }
        default:
            if (!true) {
                if (false && (((true || false) && ((-1 < 0) && false)) || (true && (127 >= v1)))) {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 uint8 = 211
                        _ = fi11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 uint8 = (^(v2))
                        if (1.5 == (-2.25 / 0.1)) {
                            v15 := GenS1{F0: -815}
                            var v16 int = (fi13 | (fi13 / ((0) | 1)))
                            _ = v15
                            _ = v16
                        } else {
                            v17 := GenS2{F0: (v8 % (((v8 | (v8 + v8))) | 1))}
                            v1 = (v1 ^ 0)
                            _ = len(v3)
                            _ = v17
                        }
                        v18 := GenS1{F0: 664}
                        _ = fi13
                        _ = v14
                        _ = v18
                    }
                }
                v8 = v8
                v8 = v8
            } else {
                switch "" {
                case "\n\t":
                    var v19 int32 = 0
                    _ = v19
                case "a":
                    var v20 uint32 = 520
                    _ = v20
                case "hello":
                    v21 := v3[(fi7 & 0)]
                    _ = v21
                }
            }
        }
        switch "" {
        case "foo bar":
            var v22 int32 = -627
            _ = v22
        case "a":
            v1 = v1
        }
        _ = fi7
        _ = v8
    }
    v1 += v1
    v1 -= 127
    var v23 uint16 = 32767
    var v24 uint16 = v23
    v25 := GenS1{F0: -9223372036854775808}
    fmt.Println(v1, v2, v24)
    v4 = append(v4, "\n\t")
    v26 := GenS3{F0: 2147483648}
    for fi27 := 0; fi27 < 6; fi27++ {
        if true {
            v1 = -110
        } else {
            for fi28 := 0; fi28 < 6; fi28++ {
                switch ("foo bar" + ("\n\t" + "foo bar")) {
                case "\n\t":
                    var v29 int16 = 42
                    _ = v29
                case "foo bar":
                    v30 := GenS1{F0: -9223372036854775808}
                    _ = v30
                case "hello":
                    var v31 int8 = (-(v1))
                    _ = v31
                }
                if len(v4) > 0 {
                    v4[0] = ((("a" + "a") + "") + ((("" + "hello") + "a") + "foo bar"))
                }
                v1 += (v1 + (v1 | 0))
                _ = fi28
            }
            fmt.Println(v1, v2, v1)
        }
        v32 := GenS2{F0: 0}
        _ = fi27
        _ = v32
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 float32 = ((0.1 * float32((0.1 + 1e10))) * (-2.25 - 0.1))
        var v35 uint16 = (^(v23))
        switch "a" {
        case "hello":
            if false {
                if (!(0 < 42)) {
                    if len(v4) > 0 {
                        v4[0] = (("a" + "foo bar") + "")
                    }
                    var v36 uint8 = (v2 + v2)
                    _ = v36
                }
                v23 = (^(v35))
                for fi37 := 0; fi37 < 14; fi37++ {
                    _ = len(v4)
                    var v38 uint = 0
                    switch (-(fi33)) {
                    case -2147483648:
                        if true {
                            v39 := GenS3{F0: 2}
                            var v40 int8 = (v1 / ((1) | 1))
                            fmt.Println(fi33, v38, v34)
                            _ = v39
                            _ = v40
                        }
                    case -1:
                        v1 += v1
                    default:
                        var v41 int32 = 1
                        _ = v41
                    }
                    _ = fi37
                    _ = v38
                }
            }
        case "foo bar":
            v42 := GenS1{F0: int64(v2)}
            _ = v42
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    v2 += v2
    v2 = (^(v2))
    if len(v4) > 0 {
        v4[0] = ("a" + ("" + (("a" + "a") + "a")))
    }
    v24 += 1
    var v43 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
