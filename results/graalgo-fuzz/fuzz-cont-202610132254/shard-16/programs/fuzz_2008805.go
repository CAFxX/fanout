package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16, p1 uint8, p2 uint32) uint8 {
    var v5 string = "hello"
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(v5, p1, p0)
        _ = fi6
    }
    v7 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return p1
}

func main() {
    var v1 string = ("a" + ("hello" + ("hello" + "a")))
    var v2 int64 = -1
    var v3 uint16 = 42
    v4 := []int{32767, 32768, 32768}
    v2 = (v2 + v2)
    switch (v1 + (v1 + "foo bar")) {
    case "a":
        v2 -= -9223372036854775808
    case "":
        if (1 >= 11) {
            v8 := GenS1{F0: 255}
            v3 -= v3
            v1 = "a"
            _ = v8
        }
    case "\n\t":
        v1 = "a"
    }
    v9 := GenS1{F0: 42}
    fmt.Println(v1, v3, v1)
    v10 := GenS1{F0: 852}
    v11 := GenS1{F0: 4294967295}
    var v12 string = v1
    if (65535 < v3) {
        switch 32767 {
        case -1:
            var v13 float64 = 0.1
            _ = v13
        case 570:
            v12 = v1
        case -585:
            v14 := GenS1{F0: 517}
            _ = v14
        }
        var v15 float32 = -2.25
        _ = v15
    }
    var v16 int64 = v2
    if (!(2147483647 >= 127)) {
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 9; fi18++ {
                switch v1 {
                case "":
                    fmt.Println(v3)
                default:
                    if true {
                        var v19 bool = (true && true)
                        _ = v19
                    }
                }
                _ = fi18
            }
            v20 := GenS1{F0: 0}
            _ = fi17
            _ = v20
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v16 += v2
        if ((v16 | (v16 << ((255) & 63))) == v2) {
            v1 = (v12 + v1)
            if true {
                fmt.Println(v1, v3, v16)
                var v21 uint32 = 225
                _ = v21
            }
            v22 := GenS1{F0: 255}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v16)
        _ = fi23
    }
    v12 = v12
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v3 = 32767
    if true {
        v1 = ""
    } else {
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 bool = (true || false)
                v16 += v2
                var v27 int64 = v2
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 int8 = -29
                v29 -= v29
                switch 128 {
                case 4294967295:
                    var v30 uint64 = 2
                    _ = v30
                case 128:
                    var v31 uint16 = v3
                    _ = v31
                case 42:
                    v1 = ("\n\t" + "a")
                default:
                    v12 = ("hello" + v1)
                }
                _ = fi28
                _ = v29
            }
            switch 292 {
            case 4294967295:
                for fi32 := 0; fi32 < 5; fi32++ {
                    var v33 float32 = (1e10 / ((1e10 * (1.5 - (0.0 + 0.0))) / -2.25))
                    v12 = v1
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi24
        }
    }
    fmt.Println(v2, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
