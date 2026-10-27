package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }
type GenS3 struct { F0 uint64; F1 string; F2 int8 }
type GenS4 struct { F0 bool; F1 uint16 }
func genFn0(p0 int16, p1 uint8, p2 uint8) string {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: (p0 >> ((2) & 63))}
        switch ("" + ("" + (("hello" + "") + "foo bar"))) {
        case "\n\t":
            if true {
                v7 := GenS1{F0: p0}
                p2 -= p1
                _ = v7
            } else {
                if (!(782 <= 2)) {
                    fmt.Println(p1, p1, p2)
                    p1 += (p1 << ((255) & 63))
                }
                v8 := []int{}
                var v9 int8 = 42
                _ = v8
                _ = v9
            }
        case "a":
            p2 = (p1 / ((p2) | 1))
        case "hello":
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 bool = false
                _ = fi10
                _ = v11
            }
        default:
            p2 = p1
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p2)
    var v12 int = -1
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := []int{fi13}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return ("\n\t" + "a")
}
func genFn1(p0 uint, p1 int8, p2 bool) bool {
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(p1, fi15)
        _ = fi15
    }
    switch 127 {
    case 32767:
        var v16 int8 = (p1 + (p1 | p1))
        _ = v16
    }
    p1 = 100
    fmt.Println(p2)
    for fi17 := 0; fi17 < 9; fi17++ {
        switch int(p1) {
        case 1:
            p2 = (p0 < (p0 * 877))
        case 32767:
            var v18 float32 = 1e10
            _ = v18
        default:
            var v19 uint = 65535
            _ = v19
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 uint = 418
    var v4 int16 = 42
    if true {
        v4 += -32768
        var v20 uint16 = 387
        var v21 int16 = 32767
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS1{F0: (v4 & (v4 << ((v3) & 63)))}
            _ = fi22
            _ = v23
        }
        if ((true || false) || false) {
            v3 += v3
        } else {
            switch 4294967295 {
            case 42:
                var v24 uint16 = 1
                _ = v24
            }
            var v25 float32 = 0.0
            _ = v25
        }
        if (!(!(false && (!(42 <= int64(v3)))))) {
            for fi26 := 0; fi26 < 2; fi26++ {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 string = "foo bar"
                    _ = len(v1)
                    v4 = v4
                    var v29 float32 = ((1.5 - 1.5) + (3.14159 - 3.14159))
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi26
            }
            v30 := v2[1]
            _ = v30
        } else {
            fmt.Println(v3, v3)
            var v31 uint16 = 255
            switch -9223372036854775808 {
            case 4294967295:
                var v32 int64 = 0
                _ = v32
            case 529:
                fmt.Println(v3, v4)
            case 32767:
                v33 := []int{386, 4294967295, 128, 128}
                _ = v33
            default:
                v3 = (v3 ^ 4294967295)
            }
            _ = v31
        }
    }
    var v34 int16 = v4
    if true {
        v35 := GenS1{F0: (v34 + 1000)}
        _ = v35
    }
    var v36 int8 = -5
    for fi37 := 0; fi37 < 3; fi37++ {
        if true {
            v4 = (v34 % (((-(v34))) | 1))
            for fi38 := 0; fi38 < 19; fi38++ {
                if (false || true) {
                    v39 := GenS2{F0: 42, F1: 0, F2: -9223372036854775808}
                    v40 := v1[1]
                    v40 -= v40
                    _ = v39
                    _ = v40
                } else {
                    var v41 int64 = -528
                    var v42 int8 = v36
                    _ = v41
                    _ = v42
                }
                _ = fi38
            }
        }
        var v43 uint16 = 237
        _ = fi37
        _ = v43
    }
    var v44 uint64 = 216
    if true {
        v36 = (^(v36))
    } else {
        v2[837] = "a"
        v45 := GenS3{F0: 95, F1: ("\n\t" + "hello"), F2: (v36 - -128)}
        _ = v45
    }
    fmt.Println(v44, v36)
    var v46 string = ("foo bar" + (((("" + "\n\t") + "\n\t") + ("a" + "foo bar")) + "\n\t"))
    for fi47 := 0; fi47 < 1; fi47++ {
        v48 := GenS4{F0: true, F1: 42}
        v49 := GenS3{F0: v44, F1: ((v46 + (v46 + ((v46 + "") + v46))) + (v46 + "")), F2: v36}
        v36 += v36
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v44
    _ = v46
}
