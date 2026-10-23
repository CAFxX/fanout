package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }

func main() {
    var v1 int8 = 127
    var v2 uint32 = 637
    var v3 int = -9223372036854775808
    var v4 int16 = -1
    var v5 int64 = -285
    var v6 uint64 = 856
    v1 = v1
    var v7 int32 = -1
    switch ("\n\t" + ("a" + ("" + "a"))) {
    case "a":
        v4 -= 1
    default:
        var v8 uint = 255
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint32 = v2
        var v11 uint16 = 2
        var v12 uint16 = (v11 | 42)
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        v1 += (v1 | v1)
        var v14 float64 = ((3.14159 + 1e10) / 3.14159)
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 42, F1: "foo bar", F2: (v5 & (v5 << ((18446744073709551615) & 63)))}
    switch ((("\n\t" + "a") + ("\n\t" + ("" + (("" + "") + ("" + ""))))) + "a") {
    case "foo bar":
        for fi16 := 0; fi16 < 7; fi16++ {
            switch (fi16 & (^(v3))) {
            case -1:
                v3 -= (v3 << ((2) & 63))
            case -2147483648:
                v17 := []int{v3, fi16}
                _ = v17
            }
            if true {
                switch 4294967295 {
                case 0:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v2 -= v2
                        _ = fi18
                    }
                case -341:
                    v4 -= v4
                case 9223372036854775807:
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v7 = v7
                        _ = fi19
                    }
                default:
                    v4 = 32767
                }
                var v20 string = "a"
                if true {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 int16 = int16(v1)
                        var v23 int16 = v22
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(v6, v1, v20)
                } else {
                    var v24 uint16 = 42
                    _ = v24
                }
                _ = v20
            } else {
                v25 := []int{}
                _ = v25
            }
            _ = fi16
        }
    case "a":
        var v26 int8 = (v1 + 0)
        _ = v26
    case "":
        fmt.Println(v6, v3, v3)
    default:
        v2 = (^(v2))
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        switch v3 {
        case 32767:
            if ((-2147483648 > v7) || false) {
                v28 := []int{fi27, v3, (v3 + v3)}
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                v28 = append(v28, v3)
                _ = v28
            } else {
                v29 := []int{fi27}
                _ = v29
            }
        case -132:
            fmt.Println(v3, v5, v3)
        case 1:
            v30 := GenS1{F0: 42, F1: ("foo bar" + (("\n\t" + "foo bar") + "hello")), F2: v5}
            _ = v30
        }
        for fi31 := 0; fi31 < 1; fi31++ {
            v6 += (v6 / (((v6 - 42)) | 1))
            var v32 int16 = v4
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            v3 -= int(v6)
            _ = fi33
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
