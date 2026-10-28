package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 string; F2 int16 }
type GenS3 struct { F0 int }
func genFn0() uint {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 65535
        for fi9 := 0; fi9 < 13; fi9++ {
            v8 = v8
            var v10 uint8 = 128
            if true {
                var v11 uint8 = v10
                fmt.Println(fi9, v8)
                fmt.Println(v8)
                switch fi7 {
                case 31:
                    var v12 int = (fi7 << ((65535) & 63))
                    _ = v12
                case 938:
                    fmt.Println(fi9)
                }
                _ = v11
            }
            v13 := GenS1{F0: (fi9 & (^(fi9)))}
            _ = fi9
            _ = v10
            _ = v13
        }
        v14 := []int{}
        _ = fi7
        _ = v14
        _ = v8
    }
    var v15 int = 567
    if (!(901 == 42)) {
        v16 := GenS2{F0: (^(v15)), F1: ("foo bar" + "\n\t"), F2: -1}
        var v17 string = ""
        _ = v16
        _ = v17
    }
    var v18 float32 = 1.5
    _ = v15
    _ = v18
    return 1
}
func genFn1(p0 uint32, p1 uint32, p2 uint16) uint8 {
    for fi19 := 0; fi19 < 18; fi19++ {
        p1 = 2147483647
        p2 -= p2
        _ = fi19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        fmt.Println(p1, fi20)
        _ = fi20
    }
    v21 := []int{}
    if true {
        p0 += 637
        if (430 == 0) {
            if ((^(p2)) == (p2 % (((^(p2))) | 1))) {
                var v22 int32 = 1
                switch "\n\t" {
                case "a":
                    v23 := GenS1{F0: -1}
                    _ = v23
                case "hello":
                    var v24 string = (("\n\t" + ("foo bar" + ("\n\t" + "a"))) + "foo bar")
                    _ = v24
                default:
                    p0 = (p0 / ((p1) | 1))
                }
                var v25 int = 32767
                _ = v22
                _ = v25
            }
            if len(v21) > 0 {
                v21[0] = 128
            }
            switch 388 {
            case -2147483648:
                var v26 int32 = 836
                _ = v26
            case 127:
                v27 := GenS1{F0: 4294967295}
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 2; fi28++ {
                p2 = 0
                _ = fi28
            }
            fmt.Println(p0, p1)
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            for fi30 := 0; fi30 < 6; fi30++ {
                if (42 <= 1000) {
                    v31 := GenS2{F0: 9223372036854775807, F1: ("\n\t" + ""), F2: 1}
                    _ = v31
                }
                p1 = p0
                var v32 int64 = -1
                p0 -= p1
                _ = fi30
                _ = v32
            }
            _ = fi29
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 127
}

func main() {
    var v1 bool = false
    var v2 uint = 65535
    var v3 bool = ((-978 >= 1) || false)
    var v4 int8 = 1
    v5 := []uint32{1, 2147483648}
    var v6 int8 = v4
    _ = len(v5)
    for fi33 := 0; fi33 < 20; fi33++ {
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 uint16 = 65535
            fmt.Println(v35, v1, v1)
            fmt.Println(v1, v4)
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    if (789 == 2147483647) {
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 bool = v3
            switch "\n\t" {
            case "\n\t":
                var v38 uint8 = 7
                _ = v38
            case "hello":
                var v39 uint = 42
                _ = v39
            default:
                v3 = (0.1 > 0.1)
            }
            _ = fi36
            _ = v37
        }
        if len(v5) > 0 {
            v5[0] = 933
        }
    } else {
        v2 += genFn0()
    }
    v6 = (v4 * v6)
    v2 = genFn0()
    var v40 uint16 = 35
    var v41 uint16 = 32767
    var v42 uint16 = 255
    v6 -= v6
    for fi43 := 0; fi43 < 9; fi43++ {
        var v44 uint32 = 42
        _ = fi43
        _ = v44
    }
    var v45 int16 = -1
    v40 = (v41 - (v40 + v40))
    if v1 {
        v1 = (2 == 127)
        for fi46 := 0; fi46 < 9; fi46++ {
            var v47 int64 = 42
            _ = fi46
            _ = v47
        }
    }
    fmt.Println(v6, v45, v42)
    var v48 uint16 = v40
    v49 := GenS3{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
