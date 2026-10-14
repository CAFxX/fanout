package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = -1
    var v2 uint = 255
    var v3 uint8 = 85
    v4 := []uint32{42, 4294967295, 32}
    v2 = (v2 + 0)
    for fi5 := 0; fi5 < 19; fi5++ {
        v2 -= (^(v2))
        _ = fi5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        if (("" + ("hello" + ("" + "foo bar"))) >= "a") {
            fmt.Println(v2)
            switch -737 {
            case 9223372036854775807:
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 int = fi7
                    v8 -= 127
                    _ = fi7
                    _ = v8
                }
            case 2147483647:
                for fi9 := 0; fi9 < 3; fi9++ {
                    fmt.Println(fi6)
                    if (true && (-2.25 > -2.25)) {
                        switch "" {
                        case "foo bar":
                            var v10 uint16 = 255
                            _ = v10
                        }
                        if len(v4) > 0 {
                            v4[0] = 2147483648
                        }
                        var v11 bool = true
                        _ = v11
                    }
                    if false {
                        switch fi9 {
                        case -9223372036854775808:
                            if true {
                                fmt.Println(fi9, v3, v3)
                                v12 := GenS1{F0: (v3 + v3)}
                                _ = v12
                            } else {
                                v1 -= (v1 << ((v2) & 63))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        case 688:
                            var v13 uint16 = 1
                            _ = v13
                        case 9223372036854775807:
                            var v14 uint32 = 453
                            _ = v14
                        }
                        var v15 int32 = 42
                        v2 = (v2 - v2)
                        _ = v15
                    } else {
                        v4 = append(v4, 2)
                        v16 := GenS1{F0: (^(v3))}
                        _ = v16
                    }
                    _ = fi9
                }
            }
            fmt.Println(fi6)
        } else {
            v4 = append(v4, 2)
            _ = len(v4)
            v4 = append(v4, 4294967295)
        }
        _ = fi6
    }
    v1 = v1
    v4 = append(v4, 2147483647)
    fmt.Println(v3, v2)
    fmt.Println(v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch ("foo bar" + "\n\t") {
    case "":
        var v17 uint = v2
        _ = v17
    case "a":
        switch 32768 {
        case 142:
            switch "" {
            case "\n\t":
                if (106 <= 700) {
                    fmt.Println(v3, v2)
                    var v18 int16 = 753
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                    _ = v18
                }
            case "hello":
                v19 := GenS1{F0: (v3 + 0)}
                _ = v19
            case "":
                fmt.Println(v2, v1, v3)
            }
        case -1:
            v20 := GenS1{F0: (^(v3))}
            _ = v20
        default:
            v3 += (v3 + (v3 << ((0) & 63)))
        }
    default:
        if false {
            switch "foo bar" {
            case "\n\t":
                v21 := GenS1{F0: 42}
                _ = v21
            case "":
                var v22 uint32 = 1
                _ = v22
            }
            var v23 string = ("\n\t" + ("a" + ("foo bar" + "")))
            var v24 int64 = v1
            _ = v23
            _ = v24
        }
    }
    v1 -= -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
