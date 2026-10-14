package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint64 = 1
    v2 := []uint32{631}
    var v3 bool = false
    v4 := make(map[string]string)
    var v5 uint8 = 255
    var v6 float32 = 1.5
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint64 = v1
        fmt.Println(v6)
        _ = fi7
        _ = v8
    }
    v3 = (v1 < v1)
    switch -2147483648 {
    case 32768:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 1:
        if (true && v3) {
            switch 4294967295 {
            case -1:
                if true {
                    v9 := GenS1{F0: -9223372036854775808}
                    _ = v9
                } else {
                    var v10 int32 = -1
                    fmt.Println(v5)
                    _ = v10
                }
            case 2147483647:
                var v11 int64 = 0
                _ = v11
            case 42:
                switch -2147483648 {
                case 127:
                    v1 = (v1 % ((v1) | 1))
                case 14:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v2 = append(v2, 2)
                        _ = fi12
                    }
                }
            }
            v13 := GenS1{F0: 513}
            _ = v13
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v6 -= (1.5 + ((-2.25 * v6) / (v6 + v6)))
            for fi14 := 0; fi14 < 15; fi14++ {
                fmt.Println(v5, fi14, v3)
                _ = fi14
            }
        }
    }
    v6 += v6
    switch 128 {
    case -9223372036854775808:
        var v15 bool = v3
        _ = v15
    default:
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint16 = 226
            fmt.Println(fi16)
            v18 := GenS1{F0: 42}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    if (!v3) {
        var v19 uint8 = 0
        _ = v19
    } else {
        var v20 int64 = 0
        var v21 uint32 = 65535
        _ = v20
        _ = v21
    }
    v1 += (v1 % (((v1 / (((v1 << ((958) & 63))) | 1))) | 1))
    v3 = (((v5 / ((167) | 1)) >= v5) || (255 != 255))
    if true {
        v22 := GenS1{F0: 42}
        _ = v22
    } else {
        v23 := GenS1{F0: 839}
        var v24 uint32 = 2147483648
        _ = v23
        _ = v24
    }
    delete(v4, "foo bar")
    v5 -= v5
    fmt.Println(v1, v5)
    var v25 uint32 = 65535
    v5 = 0
    var v26 uint8 = v5
    if (!true) {
        v25 = v25
        v5 = v26
    }
    fmt.Println(v5, v6)
    switch "\n\t" {
    case "\n\t":
        if false {
            if v3 {
                var v27 float32 = v6
                v28 := GenS2{F0: 42}
                v27 = 0.1
                _ = v27
                _ = v28
            }
            var v29 int32 = 2147483647
            if true {
                fmt.Println(v1, v1, v26)
                if v3 {
                    v30 := v4["a"]
                    fmt.Println(v30)
                    _ = v30
                } else {
                    v6 -= (v6 * (1.5 + v6))
                    v31 := GenS1{F0: -799}
                    v6 = -2.25
                    _ = v31
                }
            }
            _ = v29
        } else {
            v1 = v1
        }
    case "a":
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 int8 = 0
            v34 := GenS2{F0: 255}
            var v35 int32 = -142
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
