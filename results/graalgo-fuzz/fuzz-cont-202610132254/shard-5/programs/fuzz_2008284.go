package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []string{"\n\t", (("hello" + "hello") + "foo bar"), ((("hello" + ("\n\t" + "")) + "hello") + "")}
    v2 := []int{9223372036854775807, 42}
    v3 := []string{}
    var v4 float32 = (0.0 * ((3.14159 - (float32(1.5) - float32((0.0 + 1.5)))) + 1.5))
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 18; fi6++ {
        switch (fi6 / (((fi6 >> ((1) & 63))) | 1)) {
        case 128:
            switch fi6 {
            case -320:
                v2 = append(v2, fi6)
            case -9223372036854775808:
                v5 = (v5 * v5)
            }
        default:
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
        if (!false) {
            if ((-152 < -9223372036854775808) && (fi6 >= -2147483648)) {
                v3 = append(v3, ("a" + ((((("foo bar" + (("\n\t" + "hello") + (("a" + "") + "\n\t"))) + "foo bar") + ("" + (("" + "hello") + ("" + (("a" + "\n\t") + "foo bar"))))) + "a") + "a")))
            } else {
                if (1e10 > 0.0) {
                    if len(v2) > 0 {
                        v2[0] = 127
                    }
                    v7 := GenS1{F0: 32767}
                    _ = v7
                }
                v1 = append(v1, "")
            }
            var v8 uint32 = 2147483648
            _ = v8
        } else {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int16 = 32767
                _ = fi9
                _ = v10
            }
            switch "\n\t" {
            case "hello":
                fmt.Println(fi6, fi6)
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v4 -= (v4 / v4)
        }
        _ = fi6
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 string = "foo bar"
    if (32767 <= 839) {
        fmt.Println(v11)
    } else {
        var v12 int8 = -100
        v4 += v4
        _ = v12
    }
    switch 42 {
    case 335:
        v13 := GenS1{F0: 255}
        _ = v13
    default:
        v14 := GenS1{F0: 32767}
        _ = v14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v11 = (v11 + "a")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v4 += 3.14159
        _ = fi15
    }
    var v16 uint8 = v5
    for fi17 := 0; fi17 < 9; fi17++ {
        if (0.1 < 1.5) {
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 uint = 65535
                _ = fi18
                _ = v19
            }
            switch v11 {
            case "":
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + ("\n\t" + ("a" + v11)))
                }
            }
            v16 = (^(v16))
        } else {
            v16 -= (v5 % (((v5 % (((v5 | v5)) | 1))) | 1))
            var v20 int64 = -572
            _ = v20
        }
        fmt.Println(v11, v16)
        _ = fi17
    }
    v16 = (v5 >> ((703) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
