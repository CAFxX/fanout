package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }

func main() {
    var v1 int32 = -486
    var v2 float64 = 0.1
    v3 := []string{("\n\t" + (("a" + ("a" + "")) + ""))}
    var v4 int16 = 1000
    var v5 uint8 = 2
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint32 = 2147483647
        v9 := GenS1{F0: ("\n\t" + "foo bar"), F1: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    v5 += 1
    switch -104 {
    case 32768:
        v6[(("\n\t" + "") + (((("foo bar" + (("hello" + ("hello" + (("" + "\n\t") + "foo bar"))) + "\n\t")) + "foo bar") + "\n\t") + "foo bar"))] = 833
    default:
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int32 = (v1 << ((255) & 63))
            var v12 int = fi10
            var v13 int16 = (v4 | (v4 % ((int16(v11)) | 1)))
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    var v14 uint = 1
    if false {
        v15 := GenS1{F0: "foo bar", F1: 32767}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 19; fi16++ {
            if ((^(v14)) != v14) {
                fmt.Println(v4, v4)
            } else {
                fmt.Println(fi16)
                if ((((v1 | v1) < v1) && true) && false) {
                    v17 := GenS1{F0: "a", F1: 32767}
                    for fi18 := 0; fi18 < 17; fi18++ {
                        v14 += 2
                        if (!(float32(float64(0.1)) > -2.25)) {
                            v1 = 683
                        } else {
                            v6[("foo bar" + (("a" + "") + "a"))] = (fi16 / ((fi16) | 1))
                            var v19 uint8 = 128
                            _ = v19
                        }
                        _ = fi18
                    }
                    _ = v17
                } else {
                    var v20 uint64 = 220
                    _ = v20
                }
                var v21 bool = false
                _ = v21
            }
            _ = fi16
        }
        var v22 bool = (3.14159 == 3.14159)
        _ = v22
    }
    v4 += v4
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 int16 = v4
        var v25 int8 = -1
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: true, F1: 42, F2: 42}
    fmt.Println(v2, v14, v14)
    v14 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
