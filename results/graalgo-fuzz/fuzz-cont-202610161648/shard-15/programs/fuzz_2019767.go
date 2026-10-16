package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 int16 = 7
        _ = fi4
        _ = v5
    }
    if (!(-40 > 51)) {
        v6 := GenS1{F0: 1}
        switch 897 {
        case 2147483647:
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 float64 = -2.25
                v9 := GenS1{F0: 42}
                var v10 uint16 = 42
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        }
        v11 := []int{}
        _ = v11
        _ = v6
    } else {
        v12 := v3[678]
        _ = v12
    }
    v1 = v1
    if false {
        if (42 >= 255) {
            if (42 != 1) {
                var v13 float32 = float32(float64((((3.14159 / (1e10 * 1.5)) * 0.1) + 3.14159)))
                _ = v13
            }
            switch 906 {
            case 0:
                fmt.Println(v1, v1)
            case -2147483648:
                v14 := GenS1{F0: 429}
                _ = v14
            }
        } else {
            fmt.Println(v2, v2)
        }
    } else {
        fmt.Println(v2, v1)
    }
    _ = len(v3)
    fmt.Println(v1, v2)
    v3[2] = ""
    v2 = (v2 & v2)
    var v15 uint8 = 127
    if false {
        v1 = v1
        v16 := GenS1{F0: 155}
        fmt.Println(v15, v1, v15)
        _ = v16
    }
    v17 := GenS1{F0: 1}
    v2 -= (^(v2))
    if true {
        v2 += 832
        v18 := []int{32767, 32767}
        fmt.Println(v2)
        _ = v18
    } else {
        v19 := GenS2{F0: -1}
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v1)
            _ = fi20
        }
        _ = v19
    }
    v1 = v1
    var v21 int = -2147483648
    v22 := []int{(^(v21))}
    v2 -= (v2 & 523)
    switch ("a" + v1) {
    case "a":
        for fi23 := 0; fi23 < 15; fi23++ {
            if len(v22) > 0 {
                v22[0] = (fi23 << ((2) & 63))
            }
            fmt.Println(v21, v1, v2)
            var v24 int16 = 42
            _ = fi23
            _ = v24
        }
    case "\n\t":
        for fi25 := 0; fi25 < 5; fi25++ {
            for fi26 := 0; fi26 < 9; fi26++ {
                if true {
                    if ((v21 << ((0) & 63)) != 4294967295) {
                        var v27 int64 = 1
                        _ = v27
                    }
                } else {
                    v28 := GenS2{F0: 142}
                    if len(v22) > 0 {
                        v22[0] = (v21 >> ((23) & 63))
                    }
                    v1 = (v1 + ("\n\t" + v1))
                    v29 := GenS1{F0: 9223372036854775807}
                    if false {
                        fmt.Println(fi26, v15)
                        var v30 bool = (0 >= 572)
                        var v31 float64 = 3.14159
                        _ = v30
                        _ = v31
                    } else {
                        v32 := GenS2{F0: 0}
                        switch v21 {
                        case -2147483648:
                            var v33 float64 = float64(-2.25)
                            _ = v33
                        }
                        var v34 float64 = -2.25
                        _ = v32
                        _ = v34
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi26
            }
            _ = fi25
        }
    case "":
        var v35 uint8 = (^(v15))
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
