package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 uint64 = 32
    v5 := make(map[string]bool)
    var v6 uint16 = 469
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint = 1
        fmt.Println(v8, v6)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v1 = 0
        _ = fi9
    }
    v6 = 1
    var v10 uint16 = (v6 - 141)
    v4 = 4294967295
    v1 = v1
    v11 := v3[645]
    v6 = v6
    if true {
        fmt.Println(v10, v10)
        if (4294967295 == int(v1)) {
            var v12 float64 = 0.0
            if (!(v6 != uint16(v4))) {
                v13 := GenS1{F0: 2147483647}
                for fi14 := 0; fi14 < 12; fi14++ {
                    if (false && false) {
                        var v15 uint16 = v6
                        switch (fi14 - fi14) {
                        case 32767:
                            v4 -= v4
                        }
                        if (!(729 == 4294967295)) {
                            v10 = v10
                            switch v11 {
                            case 42:
                                for fi16 := 0; fi16 < 4; fi16++ {
                                    var v17 int64 = 1
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            var v18 uint8 = (^(v1))
                            _ = v18
                        } else {
                            v2[("\n\t" + "\n\t")] = (false || (1 >= -958))
                            var v19 int16 = -32768
                            v20 := GenS2{F0: v4}
                            _ = v19
                            _ = v20
                        }
                        _ = v15
                    }
                    _ = fi14
                }
                _ = v13
            }
            if (("\n\t" >= ((("hello" + "foo bar") + "a") + ("a" + "foo bar"))) && false) {
                v1 += v1
                v12 += 1.5
                fmt.Println(v1)
            }
            _ = v12
        }
    }
    fmt.Println(v4)
    v6 -= (v10 << ((649) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
