package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint; F1 int8; F2 uint8 }

func main() {
    v1 := []int{}
    var v2 int32 = 1
    var v3 uint16 = 2
    if true {
        v3 = 255
        if true {
            if false {
                for fi4 := 0; fi4 < 18; fi4++ {
                    v3 -= (v3 % ((v3) | 1))
                    _ = fi4
                }
                for fi5 := 0; fi5 < 6; fi5++ {
                    v2 = v2
                    _ = fi5
                }
                for fi6 := 0; fi6 < 6; fi6++ {
                    var v7 int16 = 1000
                    v8 := GenS1{F0: fi6, F1: 494}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            } else {
                var v9 uint32 = 2147483648
                switch ("" + "") {
                case "a":
                    fmt.Println(v9, v9)
                }
                fmt.Println(v3, v9)
                _ = v9
            }
            if true {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int32 = v2
                    v1 = append(v1, fi10)
                    _ = fi10
                    _ = v11
                }
                switch 4294967295 {
                case 32768:
                    v12 := GenS1{F0: -9223372036854775808, F1: 32767}
                    _ = v12
                case 127:
                    var v13 int = 534
                    _ = v13
                case 4294967295:
                    fmt.Println(v2, v2)
                }
            } else {
                v3 = 507
            }
        } else {
            v3 -= 255
            v1 = append(v1, -523)
            fmt.Println(v2, v3)
        }
        fmt.Println(v3, v3)
    } else {
        if (0.0 != 0.0) {
            var v14 uint16 = (v3 % ((v3) | 1))
            v3 = v3
            var v15 string = ("" + (("foo bar" + "a") + "hello"))
            _ = v14
            _ = v15
        }
    }
    v3 -= v3
    if ((v2 ^ v2) <= v2) {
        for fi16 := 0; fi16 < 1; fi16++ {
            v2 += (v2 % (((v2 + (^(v2)))) | 1))
            var v17 uint = 677
            v2 = v2
            _ = fi16
            _ = v17
        }
        v1 = append(v1, -1)
        for fi18 := 0; fi18 < 7; fi18++ {
            v1 = append(v1, (fi18 / ((fi18) | 1)))
            var v19 float32 = 3.14159
            _ = len(v1)
            _ = fi18
            _ = v19
        }
    } else {
        v3 += (v3 >> ((0) & 63))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            if false {
                for fi21 := 0; fi21 < 5; fi21++ {
                    if (!true) {
                        var v22 string = ""
                        v2 = -2147483648
                        var v23 int8 = 40
                        _ = v22
                        _ = v23
                    }
                    v1 = append(v1, fi21)
                    _ = fi21
                }
            }
            var v24 uint8 = 128
            _ = fi20
            _ = v24
        }
    }
    v25 := GenS2{F0: 42, F1: 42, F2: 0}
    v3 = 255
    v3 = (v3 / ((65535) | 1))
    var v26 int = 2147483647
    var v27 int64 = 9223372036854775807
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v27 = (v27 % (((v27 | -9223372036854775808)) | 1))
    switch (v26 | v26) {
    case 127:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 1:
        switch v26 {
        case -9223372036854775808:
            v27 = v27
        case -2147483648:
            var v28 string = ("hello" + "foo bar")
            _ = v28
        case 5:
            var v29 uint8 = 42
            _ = v29
        default:
            var v30 bool = (("\n\t" <= "a") && ((v2 << ((716) & 63)) != v2))
            _ = v30
        }
    case 396:
        switch "hello" {
        case "\n\t":
            var v31 int32 = v2
            _ = v31
        case "foo bar":
            _ = len(v1)
        case "":
            v32 := GenS1{F0: int(v3), F1: (v26 % ((4294967295) | 1))}
            _ = v32
        }
    }
    fmt.Println(v2, v27, v26)
    if false {
        if true {
            switch (v26 + (v26 / ((v26) | 1))) {
            case -526:
                v33 := GenS3{F0: 2147483647, F1: "foo bar"}
                _ = v33
            case 32768:
                _ = len(v1)
            }
            v3 += 0
        } else {
            v34 := GenS4{F0: 746, F1: 0, F2: 127}
            _ = v34
        }
        v27 -= (^(v27))
    }
    v35 := GenS3{F0: 4294967295, F1: "foo bar"}
    v2 = v2
    if len(v1) > 0 {
        v1[0] = (v26 * v26)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v35
}
