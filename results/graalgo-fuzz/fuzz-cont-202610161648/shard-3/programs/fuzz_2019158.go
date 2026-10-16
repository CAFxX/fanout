package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint; F2 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 bool; F1 int16; F2 int16 }
type GenS6 struct { F0 bool; F1 uint32 }
func genFn0(p0 int, p1 string, p2 int32) int64 {
    fmt.Println(p2, p2)
    p0 += -250
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() uint16 {
    var v5 uint8 = 42
    var v6 uint64 = 252
    var v7 bool = true
    if v7 {
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                v7 = (((fi8 + fi9) != 128) && true)
                v10 := GenS1{F0: 16, F1: 38}
                switch (fi8 / (((fi9 * -9223372036854775808)) | 1)) {
                case 42:
                    v11 := GenS1{F0: 0, F1: 81}
                    _ = v11
                case 1:
                    v5 += (v5 ^ v5)
                case 402:
                    v7 = false
                }
                _ = fi9
                _ = v10
            }
            v12 := GenS1{F0: 156, F1: -1}
            v13 := []int{fi8}
            _ = fi8
            _ = v12
            _ = v13
        }
    } else {
        if (42 == 1) {
            if v7 {
                for fi14 := 0; fi14 < 11; fi14++ {
                    fmt.Println(v7)
                    fmt.Println(v5, v6, v6)
                    _ = fi14
                }
                v15 := GenS2{F0: v6, F1: 127}
                _ = v15
            }
        }
        v7 = v7
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := []int{}
        fmt.Println(v5, v7)
        _ = fi16
        _ = v17
    }
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 bool = true
    var v2 uint16 = 65535
    var v3 int32 = -1
    v4 := []string{}
    for fi18 := 0; fi18 < 18; fi18++ {
        v2 -= (v2 + genFn1())
        v19 := GenS3{F0: 42, F1: 255, F2: true}
        v2 -= 32767
        _ = fi18
        _ = v19
    }
    v1 = ("a" > "foo bar")
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = ""
    }
    if v1 {
        fmt.Println(v3, v2)
    } else {
        var v20 uint16 = (v2 * v2)
        _ = v20
    }
    fmt.Println(v2, v2)
    fmt.Println(v1, v2)
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v21 := GenS1{F0: genFn1(), F1: 127}
    v3 -= 408
    v2 = (v2 + 42)
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 uint32 = 2147483647
        v24 := GenS2{F0: 18446744073709551615, F1: 1}
        var v25 int = (fi22 - fi22)
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v26 := GenS4{F0: 4294967295, F1: 42}
    if (!v1) {
        switch 1 {
        case -735:
            v27 := GenS1{F0: (v2 / ((v2) | 1)), F1: 127}
            _ = v27
        case -697:
            var v28 uint8 = 42
            _ = v28
        case 4294967295:
            switch "hello" {
            case "":
                if false {
                    fmt.Println(v2, v3)
                    if ((2147483647 != 2) || v1) {
                        v4 = append(v4, ("a" + (("" + "\n\t") + ((("hello" + ("\n\t" + ("foo bar" + ""))) + "") + ("\n\t" + ("hello" + ""))))))
                    }
                } else {
                    v29 := GenS2{F0: 723, F1: 42}
                    var v30 string = ("hello" + (("foo bar" + "foo bar") + ("hello" + ("\n\t" + ((("a" + ((("\n\t" + "a") + (("foo bar" + ("\n\t" + "foo bar")) + "hello")) + "\n\t")) + "hello") + ("" + "\n\t"))))))
                    _ = v29
                    _ = v30
                }
            case "hello":
                v1 = false
            case "a":
                var v31 int64 = -1
                _ = v31
            default:
                var v32 uint64 = 42
                _ = v32
            }
        }
        if ((-2.25 == (3.14159 + 3.14159)) || (!false)) {
            v2 -= 0
            if v1 {
                if (1 > 364) {
                    v33 := GenS5{F0: ("\n\t" >= ""), F1: 156, F2: -988}
                    v2 -= (v2 & v2)
                    _ = v33
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v3)
                }
            }
        } else {
            var v34 uint32 = 0
            var v35 int8 = 0
            v36 := GenS6{F0: (!(v1 && v1)), F1: 4294967295}
            _ = v34
            _ = v35
            _ = v36
        }
    }
    fmt.Println(v2, v3, v1)
    var v37 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v37
    _ = v4
}
