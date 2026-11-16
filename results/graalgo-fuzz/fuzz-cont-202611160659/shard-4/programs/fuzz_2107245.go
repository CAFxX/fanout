package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
func genFn0() bool {
    v5 := []int{32767, 693, 0}
    var v6 bool = (!false)
    v6 = true
    var v7 uint16 = 65535
    _ = v5
    _ = v6
    _ = v7
    return (!v6)
}
func genFn1() bool {
    v8 := GenS1{F0: ("" + "\n\t")}
    switch (((((("" + ("hello" + ("" + ("a" + "a")))) + "a") + "a") + ("" + "a")) + "") + ("" + "foo bar")) {
    case "":
        var v9 float32 = 0.0
        _ = v9
    case "a":
        for fi10 := 0; fi10 < 16; fi10++ {
            if (0.0 <= 0.0) {
                var v11 int16 = -616
                _ = v11
            } else {
                fmt.Println(fi10)
                fmt.Println(fi10)
            }
            _ = fi10
        }
    case "\n\t":
        v12 := []int{128, -1, 32767}
        _ = v12
    default:
        var v13 string = ""
        _ = v13
    }
    v14 := GenS2{F0: (1 > 128)}
    _ = v14
    _ = v8
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -219
    var v3 int = (^(v2))
    var v4 float32 = 0.0
    if false {
        var v15 int64 = -750
        fmt.Println(v4, v3)
        var v16 int32 = -65
        _ = v15
        _ = v16
    } else {
        v4 += v4
    }
    if (!(42 > -491)) {
        var v17 string = "\n\t"
        v18 := GenS2{F0: true}
        _ = v17
        _ = v18
    }
    var v19 float32 = 1e10
    v19 -= v4
    v20 := v1[339]
    switch v20 {
    case -311:
        var v21 uint8 = 1
        _ = v21
    case 879:
        _ = len(v1)
    default:
        switch v3 {
        case 2147483647:
            for fi22 := 0; fi22 < 20; fi22++ {
                v2 = (v3 % ((-1) | 1))
                switch "hello" {
                case "foo bar":
                    v2 += (v2 >> ((605) & 63))
                }
                v2 -= (^(fi22))
                _ = fi22
            }
        }
    }
    if true {
        var v23 uint32 = 2147483647
        if true {
            v23 -= (v23 & 479)
        } else {
            v24 := GenS3{F0: int32(v2)}
            _ = v24
        }
        var v25 int64 = 1
        _ = v23
        _ = v25
    } else {
        for fi26 := 0; fi26 < 14; fi26++ {
            fmt.Println(v4, v4)
            _ = fi26
        }
        if false {
            v2 = (-(v20))
            v3 = (v3 << ((4294967295) & 63))
            v20 = (v2 >> ((65535) & 63))
        } else {
            switch ("" + ((("hello" + "foo bar") + (("a" + "foo bar") + "a")) + ("foo bar" + "foo bar"))) {
            case "a":
                v20 = (v20 >> ((2) & 63))
            default:
                for fi27 := 0; fi27 < 19; fi27++ {
                    if (9223372036854775807 > (v2 ^ (v3 | v3))) {
                        if (1 != v3) {
                            var v28 float64 = 1.5
                            var v29 uint16 = 699
                            _ = v28
                            _ = v29
                        }
                    } else {
                        var v30 string = "hello"
                        _ = v30
                    }
                    v19 -= 0.0
                    if false {
                        v1[278] = (fi27 % (((^(v2))) | 1))
                        v4 -= 0.0
                        var v31 float64 = 3.14159
                        _ = v31
                    } else {
                        v1[688] = fi27
                        switch ((("\n\t" + ("" + "\n\t")) + "foo bar") + ("\n\t" + (("a" + "hello") + ("" + "\n\t")))) {
                        case "a":
                            delete(v1, 2)
                        case "":
                            if (true || ((!false) || ((v19 / v19) != float32((1.5 + 1.5))))) {
                                v32 := GenS1{F0: "hello"}
                                _ = v32
                            }
                        }
                        v1[42] = v3
                    }
                    _ = fi27
                }
            }
        }
    }
    if true {
        var v33 int64 = 0
        v34 := GenS2{F0: true}
        _ = v33
        _ = v34
    }
    fmt.Println(v19, v19)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
