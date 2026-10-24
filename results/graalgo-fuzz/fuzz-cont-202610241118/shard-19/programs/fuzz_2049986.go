package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -187
    var v3 uint = 18446744073709551615
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: 875}
    for fi8 := 0; fi8 < 15; fi8++ {
        v5["foo bar"] = ""
        v6 = (-(v6))
        _ = fi8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        if (true || true) {
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := []int{(v2 >> (((v3 ^ v3)) & 63)), 566, (-(v2)), fi10}
                if true {
                    v6 -= (-(v6))
                    v11 = append(v11, fi10)
                    v12 := GenS1{F0: 32767}
                    _ = v12
                } else {
                    var v13 uint = (v3 ^ (^(v3)))
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        }
        if ("" > (("\n\t" + (("foo bar" + "a") + ("a" + ("a" + "a")))) + ((("foo bar" + ("" + "hello")) + ("" + "a")) + "\n\t"))) {
            v6 += v6
        }
        _ = fi9
    }
    v14 := GenS1{F0: 346}
    fmt.Println(v6, v6, v2)
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            switch "a" {
            case "a":
                var v17 uint64 = 322
                _ = v17
            case "":
                v6 = 9223372036854775807
            case "\n\t":
                v6 = v6
            }
            if (!false) {
                fmt.Println(fi16, v3, v3)
                if true {
                    var v18 float64 = (3.14159 - -2.25)
                    fmt.Println(v2, fi16)
                    _ = v18
                } else {
                    if false {
                        v3 = (v3 * (v3 >> ((v3) & 63)))
                    }
                    var v19 uint8 = 0
                    if (true || true) {
                        v20 := GenS1{F0: 65535}
                        _ = v20
                    }
                    _ = v19
                }
                v3 -= v3
            }
            _ = fi16
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            fmt.Println(fi21)
            v22 := GenS1{F0: 65535}
            _ = fi21
            _ = v22
        }
        _ = fi15
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(v2, v3, fi23)
            if false {
                var v25 int8 = -1
                _ = v25
            }
            var v26 uint8 = 205
            fmt.Println(v26)
            v4["a"] = (true || (!(-1 > 1000)))
            v27 := GenS1{F0: 2}
            _ = fi24
            _ = v26
            _ = v27
        }
        _ = fi23
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 uint8 = 127
        _ = fi28
        _ = v29
    }
    var v30 float32 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
