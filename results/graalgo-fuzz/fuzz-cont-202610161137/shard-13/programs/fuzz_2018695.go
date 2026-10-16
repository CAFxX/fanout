package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 bool, p1 int, p2 uint32) string {
    var v6 int64 = 1
    v7 := []int{int(p2), 875, 32768}
    for fi8 := 0; fi8 < 3; fi8++ {
        _ = len(v7)
        p1 += fi8
        var v9 uint64 = 18446744073709551615
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 420
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    return (("foo bar" + "a") + "\n\t")
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{(^(v1))}
    var v3 string = ""
    v4 := []uint32{4294967295, (v1 | v1), (v1 & 42)}
    var v5 string = v3
    switch "\n\t" {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "":
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: 42, F1: fi11, F2: 32767}
            _ = fi11
            _ = v12
        }
    }
    if true {
        if true {
            switch v5 {
            case "hello":
                var v13 int32 = 1
                _ = v13
            case "foo bar":
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 uint = 561
                    var v16 int16 = int16(v1)
                    var v17 int32 = -122
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            default:
                var v18 int32 = 386
                _ = v18
            }
        }
        var v19 bool = (127 > -1)
        var v20 uint32 = (^(v1))
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            v3 = (v3 + (v5 + (v5 + "foo bar")))
            v1 -= v1
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 uint8 = 1
                var v24 uint8 = 26
                switch v5 {
                case "foo bar":
                    v23 = v24
                    v1 = 2
                default:
                    v1 = v1
                    var v25 bool = (-1 != 528)
                    var v26 int8 = 42
                    _ = v25
                    _ = v26
                }
                fmt.Println(fi22)
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        if len(v4) > 0 {
            v4[0] = v1
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            _ = len(v4)
            if true {
                for fi28 := 0; fi28 < 4; fi28++ {
                    v2 = append(v2, v1)
                    v29 := GenS2{F0: -128, F1: 2, F2: 102}
                    var v30 uint64 = 2
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            } else {
                if ((32767 != 42) || (469 < -2147483648)) {
                    switch (fi27 & fi27) {
                    case 127:
                        var v31 uint16 = 425
                        _ = v31
                    default:
                        var v32 int32 = 0
                        _ = v32
                    }
                }
            }
            for fi33 := 0; fi33 < 14; fi33++ {
                var v34 int = (fi33 + fi27)
                _ = fi33
                _ = v34
            }
            _ = fi27
        }
    }
    v35 := GenS3{F0: 233, F1: 42}
    v3 = v3
    for fi36 := 0; fi36 < 3; fi36++ {
        v5 = (((v5 + "\n\t") + ((v3 + v5) + ((v5 + "\n\t") + "\n\t"))) + v5)
        _ = fi36
    }
    fmt.Println(v5)
    if true {
        for fi37 := 0; fi37 < 16; fi37++ {
            var v38 float64 = 1e10
            _ = fi37
            _ = v38
        }
    }
    if false {
        v39 := GenS4{F0: ((v5 + v3) + ((((((v5 + v5) + (v5 + "")) + ("\n\t" + ("hello" + v3))) + ("hello" + ("hello" + v3))) + v3) + ("a" + "foo bar"))), F1: (1 != 197)}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
