package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 bool; F1 int64 }
func genFn0(p0 int32, p1 int16) uint64 {
    p1 += 1
    v7 := []int{32768, 871, 0}
    v7 = append(v7, -1)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int = 128
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if (false || true) {
                fmt.Println(fi11, p1)
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v12 uint8 = 127
                _ = v12
            }
            v13 := GenS1{F0: p1, F1: 5}
            fmt.Println(p0, p1)
            var v14 int8 = 127
            _ = fi11
            _ = v13
            _ = v14
        }
        p1 = (p1 >> ((564) & 63))
        var v15 uint32 = 2147483647
        _ = fi10
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint = 167
    v2 := make(map[uint32]bool)
    v3 := []string{("\n\t" + (("" + "\n\t") + (("\n\t" + ((("foo bar" + "a") + ("a" + (("a" + (("" + "a") + ("a" + "foo bar"))) + ""))) + "\n\t")) + (("foo bar" + ("a" + "foo bar")) + "foo bar")))), "hello", "", "a"}
    v4 := []uint32{2147483648}
    var v5 int8 = 36
    var v6 uint64 = 2
    if (true && (-921 != 422)) {
        switch "foo bar" {
        case "\n\t":
            for fi16 := 0; fi16 < 3; fi16++ {
                for fi17 := 0; fi17 < 12; fi17++ {
                    v18 := GenS1{F0: -32768, F1: (v1 >> ((v1) & 63))}
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        case "hello":
            v19 := v2[2]
            _ = v19
        case "":
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 string = "a"
                v22 := GenS2{F0: v1, F1: 42, F2: true}
                _ = fi20
                _ = v21
                _ = v22
            }
        default:
            var v23 string = "hello"
            _ = v23
        }
    } else {
        if (!true) {
            v24 := GenS2{F0: 124, F1: 42, F2: (970 > 32767)}
            _ = v24
        } else {
            fmt.Println(v6, v1, v6)
            switch ((("\n\t" + ("\n\t" + "foo bar")) + ("\n\t" + "foo bar")) + "\n\t") {
            case "foo bar":
                _ = len(v3)
            case "":
                delete(v2, 2147483648)
            case "\n\t":
                for fi25 := 0; fi25 < 14; fi25++ {
                    fmt.Println(v6, fi25, v6)
                    if false {
                        v6 -= genFn0(42, 0)
                    }
                    var v26 int16 = -74
                    _ = fi25
                    _ = v26
                }
            }
        }
        v27 := GenS3{F0: v1, F1: -1}
        _ = v27
    }
    fmt.Println(v6, v5, v1)
    v3 = append(v3, ("hello" + "a"))
    v28 := GenS1{F0: -32768, F1: v1}
    v5 = (v5 - int8(v6))
    var v29 uint = 255
    switch ("a" + "\n\t") {
    case "hello":
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int16 = 1
            var v32 uint = v29
            _ = fi30
            _ = v31
            _ = v32
        }
    default:
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 int = int(v1)
            v3 = append(v3, "a")
            var v35 int = (v34 ^ (fi33 - 32767))
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        var v37 uint = (v1 >> (((v1 % ((v29) | 1))) & 63))
        _ = fi36
        _ = v37
    }
    v38 := GenS4{F0: (-293 == 114), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
