package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint16; F1 int16 }

func main() {
    var v1 int8 = -128
    var v2 uint = 255
    var v3 bool = false
    var v4 string = ("hello" + "foo bar")
    var v5 int32 = -2147483648
    v1 -= (v1 % ((v1) | 1))
    for fi6 := 0; fi6 < 20; fi6++ {
        v5 += (v5 + (v5 ^ 28))
        _ = fi6
    }
    var v7 uint32 = 605
    if false {
        v8 := []int{0, 127, -1}
        _ = len(v8)
        _ = v8
    } else {
        v7 = (^(v7))
    }
    v4 = "a"
    switch 9223372036854775807 {
    case -118:
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 uint = (^(v2))
                v4 = "hello"
                v2 += 65535
                _ = fi10
                _ = v11
            }
            switch (v4 + "foo bar") {
            case "\n\t":
                var v12 uint = v2
                _ = v12
            case "hello":
                for fi13 := 0; fi13 < 6; fi13++ {
                    fmt.Println(v4, v4)
                    v14 := []int{int(v2), (fi13 / (((fi13 ^ 128)) | 1)), (fi9 & (fi9 >> ((v2) & 63))), fi9}
                    _ = fi13
                    _ = v14
                }
            default:
                var v15 int = (fi9 & (fi9 ^ fi9))
                _ = v15
            }
            _ = fi9
        }
    default:
        var v16 string = (v4 + "\n\t")
        _ = v16
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v3, v7, v7)
    case "a":
        v1 -= (v1 ^ -95)
    }
    v17 := []int{1, -9223372036854775808, 9223372036854775807}
    switch (v4 + (("foo bar" + "") + (v4 + "hello"))) {
    case "a":
        fmt.Println(v4, v4)
    case "hello":
        if (-32768 >= -1) {
            v17 = append(v17, 128)
            for fi18 := 0; fi18 < 10; fi18++ {
                for fi19 := 0; fi19 < 13; fi19++ {
                    var v20 uint32 = (v7 >> ((18446744073709551615) & 63))
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 16; fi21++ {
                    v1 += v1
                    _ = fi21
                }
                var v22 bool = v3
                _ = fi18
                _ = v22
            }
            var v23 bool = (!false)
            _ = v23
        } else {
            if (-505 > int64(v5)) {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                var v24 int32 = v5
                _ = v24
            }
            switch (v4 + (v4 + ("\n\t" + "foo bar"))) {
            case "\n\t":
                _ = len(v17)
            case "a":
                v7 = 564
            default:
                if ((v3 && v3) && v3) {
                    fmt.Println(v1)
                    v3 = true
                }
            }
        }
    }
    v25 := GenS1{F0: 4294967295, F1: v5, F2: 255}
    v26 := GenS1{F0: 661, F1: v5, F2: 1}
    for fi27 := 0; fi27 < 8; fi27++ {
        fmt.Println(v7, v7, v3)
        for fi28 := 0; fi28 < 8; fi28++ {
            _ = len(v17)
            _ = fi28
        }
        if (((("" + (v4 + v4)) + v4) + "foo bar") != ("a" + ((v4 + "hello") + "foo bar"))) {
            for fi29 := 0; fi29 < 11; fi29++ {
                fmt.Println(fi29)
                var v30 float32 = 1.5
                if (((v2 % ((18446744073709551615) | 1)) == 2) && v3) {
                    var v31 float32 = (float32(((3.14159 - float64(v30)) / (-2.25 * (0.1 + 1e10)))) / v30)
                    v32 := GenS2{F0: 0}
                    var v33 int32 = v5
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = fi29
                _ = v30
            }
        } else {
            v1 = int8(v7)
            v2 = 0
        }
        _ = fi27
    }
    v34 := GenS2{F0: v1}
    var v35 string = "hello"
    v36 := GenS3{F0: -384, F1: 0}
    _ = len(v17)
    if v3 {
        if (!(0 < 133)) {
            if (v3 && (v7 < v7)) {
                var v37 int8 = (v1 ^ v1)
                v38 := GenS4{F0: 1, F1: -32768}
                for fi39 := 0; fi39 < 19; fi39++ {
                    v4 = v35
                    _ = fi39
                }
                _ = v37
                _ = v38
            } else {
                v17 = append(v17, 127)
            }
            var v40 string = ((v35 + v4) + v4)
            for fi41 := 0; fi41 < 16; fi41++ {
                for fi42 := 0; fi42 < 11; fi42++ {
                    v4 = "\n\t"
                    var v43 int = (fi42 - fi41)
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    switch v35 {
                    case "\n\t":
                        v2 += (v2 * v2)
                    case "":
                        var v44 float32 = ((0.1 + 3.14159) + 1e10)
                        _ = v44
                    case "a":
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    default:
                        var v45 int16 = 522
                        _ = v45
                    }
                    v35 = (v4 + v40)
                    var v46 int8 = v1
                    v7 += v7
                    var v47 int32 = v5
                    _ = fi42
                    _ = v43
                    _ = v46
                    _ = v47
                }
                v4 = v4
                var v48 int8 = v1
                _ = fi41
                _ = v48
            }
            _ = v40
        } else {
            if v3 {
                fmt.Println(v4)
                fmt.Println(v1, v1, v1)
            } else {
                if true {
                    var v49 string = ("\n\t" + ("foo bar" + ((v4 + v35) + (v35 + (("" + v4) + (v35 + (("foo bar" + v4) + (v4 + "foo bar"))))))))
                    if v3 {
                        switch ((v35 + "a") + ((v35 + v4) + v49)) {
                        case "hello":
                            for fi50 := 0; fi50 < 3; fi50++ {
                                var v51 float32 = 3.14159
                                for fi52 := 0; fi52 < 11; fi52++ {
                                    v2 = (v2 >> (((^(v2))) & 63))
                                    _ = fi52
                                }
                                _ = fi50
                                _ = v51
                            }
                        case "":
                            v17 = append(v17, -1)
                        }
                        var v53 float32 = 0.1
                        _ = v53
                    }
                    if (v7 != (v7 % ((v7) | 1))) {
                        if v3 {
                            _ = len(v17)
                        } else {
                            if len(v17) > 0 {
                                v17[0] = -833
                            }
                            v1 -= (v1 + 0)
                        }
                    } else {
                        var v54 int = 521
                        v5 = v5
                        if len(v17) > 0 {
                            v17[0] = (-(v54))
                        }
                        _ = v54
                    }
                    _ = v49
                }
                v4 = "hello"
            }
            for fi55 := 0; fi55 < 12; fi55++ {
                switch (v4 + v4) {
                case "\n\t":
                    var v56 string = v35
                    _ = v56
                }
                var v57 float64 = 1.5
                fmt.Println(v2, v5)
                _ = fi55
                _ = v57
            }
        }
        var v58 uint = 1
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v7
}
