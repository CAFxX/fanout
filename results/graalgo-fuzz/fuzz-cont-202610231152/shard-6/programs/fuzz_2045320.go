package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string }

func main() {
    var v1 int16 = 347
    v2 := []int{32768}
    v3 := []string{"foo bar", "foo bar", "foo bar"}
    v1 -= v1
    var v4 int32 = 1
    if false {
        v4 += v4
    }
    if false {
        for fi5 := 0; fi5 < 9; fi5++ {
            v1 = (v1 % ((int16(v4)) | 1))
            v4 = (-(v4))
            _ = fi5
        }
        var v6 float32 = 0.0
        _ = v6
    } else {
        fmt.Println(v1, v4)
        v4 += (v4 & (v4 & v4))
    }
    if len(v2) > 0 {
        v2[0] = 32768
    }
    v4 += v4
    for fi7 := 0; fi7 < 20; fi7++ {
        v2 = append(v2, fi7)
        switch -1 {
        case 9223372036854775807:
            var v8 uint32 = 0
            _ = v8
        case -352:
            switch ("foo bar" + ("\n\t" + "")) {
            case "\n\t":
                for fi9 := 0; fi9 < 10; fi9++ {
                    v1 -= v1
                    fmt.Println(fi9, v1)
                    v4 = -2147483648
                    _ = fi9
                }
            case "":
                _ = len(v3)
            }
        case -323:
            if len(v3) > 0 {
                v3[0] = ""
            }
        default:
            switch (("hello" + "foo bar") + ("hello" + "")) {
            case "":
                v1 = int16(fi7)
            }
        }
        v10 := GenS1{F0: (("a" + "\n\t") + "foo bar"), F1: ("" + "")}
        _ = fi7
        _ = v10
    }
    var v11 uint16 = 255
    for fi12 := 0; fi12 < 9; fi12++ {
        v4 += v4
        _ = fi12
    }
    var v13 bool = true
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v14 float32 = 3.14159
    if v13 {
        if (519 == 65535) {
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint = 2
                _ = fi15
                _ = v16
            }
            v14 = v14
        } else {
            fmt.Println(v4, v1, v14)
            v14 -= ((3.14159 * float32((float64(float32(-2.25)) * float64(((1e10 - v14) * -2.25))))) / v14)
        }
        for fi17 := 0; fi17 < 8; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                if false {
                    v4 = (v4 << ((4294967295) & 63))
                    if (!(!v13)) {
                        v19 := GenS2{F0: "foo bar"}
                        if false {
                            var v20 int = fi17
                            if len(v3) > 0 {
                                v3[0] = ("\n\t" + "foo bar")
                            }
                            v4 = 42
                            _ = v20
                        }
                        _ = v19
                    } else {
                        v1 -= v1
                    }
                    var v21 int64 = 0
                    _ = v21
                }
                v22 := GenS1{F0: (("hello" + ("" + "\n\t")) + "\n\t"), F1: "foo bar"}
                _ = fi18
                _ = v22
            }
            _ = fi17
        }
        switch -2147483648 {
        case -2147483648:
            for fi23 := 0; fi23 < 1; fi23++ {
                if v13 {
                    switch 127 {
                    case 658:
                        v24 := GenS1{F0: "", F1: "foo bar"}
                        _ = v24
                    }
                }
                _ = fi23
            }
        case -1:
            var v25 int16 = (v1 * v1)
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
