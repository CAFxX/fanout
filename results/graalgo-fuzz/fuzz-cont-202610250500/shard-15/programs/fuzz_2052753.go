package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint, p1 bool, p2 string) uint64 {
    fmt.Println(p1, p2, p0)
    var v5 float32 = ((1e10 - 1.5) + ((1e10 / (3.14159 / 1.0)) / (float32((((1e10 - ((3.14159 - (((((float64(1.5) - 1.5) / 0.1) / 3.14159) - 0.1) * 0.0)) - 1.5)) / 1.0) - ((1.5 - 0.0) + -2.25))) * ((1.5 + 3.14159) - float32(3.14159)))))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return uint64(p0)
}

func main() {
    var v1 uint8 = 34
    v2 := []string{"a"}
    var v3 uint64 = 1
    var v4 float64 = 0.1
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float64 = -2.25
        switch "a" {
        case "\n\t":
            var v8 int32 = -2147483648
            _ = v8
        case "a":
            v4 -= v4
        default:
            v4 += v7
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: (v3 << ((2) & 63))}
            switch ("foo bar" + "foo bar") {
            case "a":
                _ = len(v2)
            }
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    if (1 <= -252) {
        v2 = append(v2, ("" + (("foo bar" + ("\n\t" + ((("\n\t" + "hello") + "") + ("" + "foo bar")))) + "")))
        var v11 uint = 4294967295
        switch "\n\t" {
        case "a":
            fmt.Println(v1, v4, v11)
        case "hello":
            for fi12 := 0; fi12 < 11; fi12++ {
                for fi13 := 0; fi13 < 10; fi13++ {
                    switch (^(fi12)) {
                    case 9223372036854775807:
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            fmt.Println(v11, fi13)
                        default:
                            switch fi12 {
                            case 9223372036854775807:
                                v14 := GenS1{F0: genFn0(uint(v3), (18446744073709551615 <= (v11 % (((v11 % ((v11) | 1))) | 1))), "hello")}
                                _ = v14
                            default:
                                v2 = append(v2, ("a" + ("foo bar" + "\n\t")))
                            }
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    _ = fi13
                }
                _ = fi12
            }
        case "":
            switch 4294967295 {
            case 128:
                fmt.Println(v4, v4)
            }
        }
        _ = v11
    }
    if (255 <= v1) {
        v2 = append(v2, "a")
        var v15 int64 = 1
        _ = v15
    }
    var v16 string = "a"
    fmt.Println(v4, v4)
    var v17 int64 = -726
    switch 128 {
    case 9223372036854775807:
        v1 = v1
    case 4294967295:
        v3 = (v3 << ((42) & 63))
    }
    var v18 int8 = -1
    v1 = (^(v1))
    switch (v16 + (v16 + "foo bar")) {
    case "a":
        for fi19 := 0; fi19 < 6; fi19++ {
            if (((-2.25 + (0.1 * 0.1)) * (v4 * 0.1)) <= (v4 + float64(1e10))) {
                if true {
                    var v20 int64 = v17
                    var v21 int64 = (v20 + (v17 >> ((0) & 63)))
                    for fi22 := 0; fi22 < 2; fi22++ {
                        var v23 int32 = 0
                        v20 = v17
                        _ = fi22
                        _ = v23
                    }
                    _ = v20
                    _ = v21
                } else {
                    v16 = v16
                }
            } else {
                v17 += (v17 & (v17 + (^(v17))))
                var v24 uint16 = 2
                fmt.Println(v3, v17, v4)
                _ = v24
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                if len(v2) > 0 {
                    v2[0] = v16
                }
                var v26 uint16 = 1
                _ = fi25
                _ = v26
            }
            v27 := GenS1{F0: (v3 | (v3 ^ (^(v3))))}
            _ = fi19
            _ = v27
        }
    }
    var v28 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
