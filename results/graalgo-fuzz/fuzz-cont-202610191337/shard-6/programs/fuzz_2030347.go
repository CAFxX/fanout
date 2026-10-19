package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int; F1 bool; F2 uint }

func main() {
    v1 := []uint32{2, 65535}
    var v2 uint32 = 2
    v3 := []string{(((("" + "a") + "\n\t") + ("\n\t" + "foo bar")) + "foo bar"), ("" + "hello")}
    var v4 uint32 = 2
    var v5 bool = true
    switch 128 {
    case 42:
        for fi6 := 0; fi6 < 13; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                switch (((("\n\t" + ("a" + ("\n\t" + ""))) + "hello") + "") + "a") {
                case "":
                    var v8 uint64 = 42
                    _ = v8
                default:
                    fmt.Println(fi6, v4, fi7)
                }
                var v9 uint16 = 255
                v5 = false
                _ = fi7
                _ = v9
            }
            var v10 int8 = -128
            _ = fi6
            _ = v10
        }
    case 0:
        var v11 string = ""
        _ = v11
    default:
        v2 = (v4 & 1)
    }
    if len(v3) > 0 {
        v3[0] = (("" + ("hello" + (("foo bar" + "foo bar") + ("" + "\n\t")))) + "a")
    }
    switch "foo bar" {
    case "a":
        v12 := GenS1{F0: (723 > 418), F1: false, F2: 768}
        _ = v12
    case "":
        v4 -= 0
    case "hello":
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    default:
        var v13 uint16 = 65535
        _ = v13
    }
    if (!(-532 > 230)) {
        if (18446744073709551615 != 1) {
            for fi14 := 0; fi14 < 5; fi14++ {
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                v2 -= v2
                for fi15 := 0; fi15 < 15; fi15++ {
                    v2 = (v2 - v2)
                    v5 = (!(128 >= 0))
                    _ = fi15
                }
                _ = fi14
            }
            var v16 uint64 = 4294967295
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := GenS1{F0: (1.5 == -2.25), F1: v5, F2: 255}
                _ = len(v1)
                v19 := GenS1{F0: v5, F1: true, F2: 821}
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v16
        } else {
            v2 -= 4294967295
            v20 := GenS2{F0: 32767}
            _ = v20
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 uint16 = 487
            v5 = (false || false)
            _ = fi21
            _ = v22
        }
    } else {
        var v23 int8 = -128
        _ = v23
    }
    var v24 bool = (!false)
    switch 73 {
    case 2147483647:
        v24 = (63 >= 42)
    case 9223372036854775807:
        var v25 uint = 74
        _ = v25
    default:
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 uint8 = 127
            fmt.Println(v2, v4, v2)
            v3 = append(v3, "foo bar")
            _ = fi26
            _ = v27
        }
    }
    v5 = false
    for fi28 := 0; fi28 < 10; fi28++ {
        v24 = false
        v29 := GenS1{F0: (128 > (-(fi28))), F1: v24, F2: 588}
        _ = fi28
        _ = v29
    }
    switch 9223372036854775807 {
    case 128:
        switch -362 {
        case 595:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 4294967295:
            if v24 {
                v30 := GenS3{F0: 559, F1: 130}
                for fi31 := 0; fi31 < 1; fi31++ {
                    v3 = append(v3, ("\n\t" + "foo bar"))
                    _ = fi31
                }
                _ = v30
            }
        default:
            v24 = true
        }
    default:
        v2 = (v4 << ((255) & 63))
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 bool = true
        v34 := GenS4{F0: (fi32 << ((uint(v2)) & 63)), F1: ("foo bar" == "foo bar"), F2: 372}
        if v5 {
            switch "foo bar" {
            case "a":
                v4 = 407
            case "hello":
                _ = len(v1)
            default:
                v35 := GenS4{F0: (-(fi32)), F1: (0 != 42), F2: 1}
                _ = v35
            }
            v36 := GenS2{F0: 32767}
            _ = v36
        } else {
            var v37 int64 = 1
            _ = v37
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    var v38 uint32 = 0
    if (v24 || (false || (!false))) {
        v3 = append(v3, "hello")
        v39 := GenS1{F0: (1 > 9223372036854775807), F1: v24, F2: 121}
        v2 = v4
        _ = v39
    }
    v24 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
