package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() int16 {
    switch "foo bar" {
    case "foo bar":
        v5 := []int{4294967295, -869, -1, 9223372036854775807}
        _ = v5
    case "\n\t":
        var v6 float32 = (float32(3.14159) * 1.5)
        _ = v6
    case "a":
        v7 := GenS1{F0: (false || false)}
        _ = v7
    }
    var v8 int32 = -918
    _ = v8
    return 32767
}
func genFn1(p0 uint16) string {
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(p0)
        switch -718 {
        case 546:
            switch (-(fi9)) {
            case 4294967295:
                switch (fi9 - 2147483647) {
                case 9223372036854775807:
                    if true {
                        for fi10 := 0; fi10 < 18; fi10++ {
                            p0 -= (^(p0))
                            v11 := []int{(-(fi9)), (fi10 + (^(fi10))), (fi10 + int(p0))}
                            _ = fi10
                            _ = v11
                        }
                    }
                }
            case -5:
                p0 = (p0 / ((255) | 1))
            default:
                if (!false) {
                    if true {
                        var v12 int32 = 0
                        var v13 uint8 = 2
                        var v14 int8 = -33
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v15 float64 = -2.25
                    v16 := []int{}
                    _ = v15
                    _ = v16
                } else {
                    fmt.Println(p0)
                }
            }
        case 951:
            v17 := []int{-2147483648, (fi9 << ((65535) & 63))}
            _ = v17
        default:
            p0 = 0
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            p0 += p0
            var v19 float32 = 3.14159
            _ = fi18
            _ = v19
        }
        _ = fi9
    }
    v20 := GenS1{F0: true}
    _ = p0
    _ = v20
    return ""
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := []string{""}
    v4 := make(map[string]bool)
    if true {
        var v21 float64 = (float64(0.1) / ((1.5 - -2.25) * 3.14159))
        if false {
            var v22 int = 497
            _ = v22
        } else {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v21 += 3.14159
        }
        _ = v21
    }
    v23 := GenS1{F0: false}
    var v24 uint16 = 2
    var v25 float32 = 3.14159
    _ = len(v2)
    v26 := GenS1{F0: (true || (2 != 0))}
    v24 = v24
    v3 = append(v3, (("foo bar" + "\n\t") + "hello"))
    for fi27 := 0; fi27 < 17; fi27++ {
        v24 += 255
        _ = len(v3)
        _ = fi27
    }
    var v28 uint64 = 922
    if (673 >= 44) {
        v29 := GenS1{F0: (true || (true && (255 != 42)))}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        delete(v1, 993)
        _ = v29
    } else {
        _ = len(v4)
        var v30 uint64 = (^(v28))
        var v31 int8 = 100
        _ = v30
        _ = v31
    }
    _ = len(v4)
    fmt.Println(v25)
    v25 += v25
    for fi32 := 0; fi32 < 14; fi32++ {
        v28 = v28
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 string = "foo bar"
            if ((!true) && false) {
                var v35 uint64 = 920
                _ = v35
            }
            v24 += uint16(v28)
            _ = fi33
            _ = v34
        }
        v25 = -2.25
        _ = fi32
    }
    v2[65535] = 981
    delete(v4, "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
