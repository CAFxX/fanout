package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint32) int {
    v4 := []int{-167}
    var v5 int = 32768
    for fi6 := 0; fi6 < 7; fi6++ {
        if len(v4) > 0 {
            v4[0] = (fi6 << ((18446744073709551615) & 63))
        }
        _ = fi6
    }
    fmt.Println(v5)
    p0 += p0
    _ = p0
    _ = v4
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 uint8, p1 uint64) uint16 {
    p0 -= (p0 >> ((18446744073709551615) & 63))
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{fi7}
        _ = fi7
        _ = v8
    }
    switch 0 {
    case 128:
        switch ("\n\t" + ((("\n\t" + ("hello" + (("foo bar" + ("" + "foo bar")) + "foo bar"))) + "hello") + "\n\t")) {
        case "hello":
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(p0, p1, p1)
                v10 := []int{}
                _ = fi9
                _ = v10
            }
        }
    case -1:
        v11 := []int{2147483647, 32768}
        _ = v11
    case 32768:
        v12 := []int{4294967295, 32768}
        _ = v12
    default:
        var v13 uint64 = p1
        _ = v13
    }
    switch 32768 {
    case -994:
        p1 -= 439
    default:
        p1 += 18446744073709551615
    }
    v14 := []int{128, 2147483647, 128, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v14
    return 255
}

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{uint32(v1), 4294967295, 2}
    var v3 int32 = -825
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            v1 = 65535
            if (3.14159 <= (0.1 + float64(3.14159))) {
                var v17 uint16 = v1
                _ = v17
            }
            v3 = v3
            _ = fi16
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            if true {
                switch "" {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "hello":
                    switch (("hello" + "\n\t") + "") {
                    case "foo bar":
                        v1 = 146
                    }
                default:
                    if true {
                        fmt.Println(fi18, fi15)
                    } else {
                        var v19 int = (fi15 - fi15)
                        var v20 uint = 981
                        _ = v19
                        _ = v20
                    }
                }
                v21 := GenS1{F0: v3}
                v2 = append(v2, 65535)
                var v22 uint = 236
                _ = v21
                _ = v22
            } else {
                if ("foo bar" <= "\n\t") {
                    var v23 uint8 = 119
                    if ((false && (false && false)) && true) {
                        if len(v2) > 0 {
                            v2[0] = 523
                        }
                    }
                    var v24 string = "a"
                    _ = v23
                    _ = v24
                }
            }
            v1 = (v1 & genFn1(uint8(v1), 2))
            if ((42 < -1) && (!(!false))) {
                if (2 != 474) {
                    switch 32767 {
                    case 32768:
                        var v25 string = ""
                        _ = v25
                    }
                    fmt.Println(v1)
                    var v26 int = 42
                    _ = v26
                }
            } else {
                switch (("foo bar" + "a") + "") {
                case "a":
                    _ = len(v2)
                }
                _ = len(v2)
                v3 -= (v3 | 0)
            }
            _ = fi18
        }
        _ = fi15
    }
    var v27 uint32 = 65535
    switch 1 {
    case 127:
        var v28 uint8 = 128
        _ = v28
    }
    fmt.Println(v27)
    v27 += (v27 + (v27 | 403))
    var v29 int16 = -1
    v1 = (^(v1))
    v3 = (v3 ^ (v3 << ((37) & 63)))
    v27 -= (v27 - 1)
    var v30 int = 32768
    v2 = append(v2, (v27 ^ v27))
    if true {
        v31 := GenS2{F0: 1}
        for fi32 := 0; fi32 < 6; fi32++ {
            for fi33 := 0; fi33 < 5; fi33++ {
                fmt.Println(v27, v29, v1)
                switch "hello" {
                case "hello":
                    var v34 int = int(v29)
                    _ = v34
                case "\n\t":
                    v35 := GenS1{F0: -42}
                    _ = v35
                case "foo bar":
                    fmt.Println(v1)
                default:
                    fmt.Println(v27, v29)
                }
                _ = fi33
            }
            _ = fi32
        }
        _ = v31
    }
    fmt.Println(v3, v29, v3)
    v3 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
}
