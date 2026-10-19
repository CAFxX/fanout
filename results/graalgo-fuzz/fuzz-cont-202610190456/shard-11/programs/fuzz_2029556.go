package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []int{}
    var v2 bool = true
    v3 := []uint32{65535, 2147483647, 4294967295}
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 42}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        v2 = true
        _ = fi6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v2 = (!(v2 && true))
        var v8 int8 = 45
        _ = fi7
        _ = v8
    }
    v2 = true
    v9 := GenS1{F0: 100}
    if len(v3) > 0 {
        v3[0] = 702
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if (!(!true)) {
            v2 = (!(-32768 == 1000))
        } else {
            _ = len(v1)
        }
        fmt.Println(fi10)
        _ = fi10
    }
    if v2 {
        var v11 int16 = 32767
        _ = v11
    } else {
        v2 = (446 >= 2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 = (1.5 != 0.0)
    if false {
        var v12 uint8 = 255
        if (false && ((76 > 42) && ((v12 >> ((65535) & 63)) > (v12 % ((v12) | 1))))) {
            var v13 int8 = 127
            fmt.Println(v2, v12, v2)
            _ = v13
        } else {
            if ((v2 || v2) || v2) {
                v2 = (65535 >= 4294967295)
            } else {
                if v2 {
                    if (v2 && (334 == 255)) {
                        var v14 uint32 = 467
                        switch -1 {
                        case 753:
                            v14 += v14
                        default:
                            var v15 bool = ((v2 || v2) || true)
                            _ = v15
                        }
                        _ = v14
                    } else {
                        var v16 string = "hello"
                        var v17 float32 = 1.5
                        _ = v16
                        _ = v17
                    }
                    if len(v1) > 0 {
                        v1[0] = -2147483648
                    }
                } else {
                    switch -97 {
                    case 128:
                        v12 = (v12 % (((v12 / (((v12 | (v12 + v12))) | 1))) | 1))
                    case -203:
                        for fi18 := 0; fi18 < 16; fi18++ {
                            var v19 uint16 = 790
                            var v20 string = "foo bar"
                            v20 = v20
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    case 896:
                        v2 = (!(!false))
                    }
                    switch "a" {
                    case "\n\t":
                        for fi21 := 0; fi21 < 5; fi21++ {
                            var v22 uint32 = 0
                            _ = fi21
                            _ = v22
                        }
                    case "hello":
                        for fi23 := 0; fi23 < 13; fi23++ {
                            var v24 uint16 = 222
                            v2 = (-128 <= 42)
                            v24 = (v24 - v24)
                            _ = fi23
                            _ = v24
                        }
                    case "foo bar":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        v2 = (-9223372036854775808 <= 0)
                    }
                    if (42 != 0) {
                        for fi25 := 0; fi25 < 8; fi25++ {
                            fmt.Println(v2, v2)
                            fmt.Println(v12, fi25, fi25)
                            v3 = append(v3, 2147483648)
                            _ = fi25
                        }
                        v12 -= 0
                        var v26 int16 = -32768
                        _ = v26
                    }
                }
                var v27 float32 = -2.25
                _ = v27
            }
            v1 = append(v1, 128)
        }
        _ = v12
    }
    var v28 uint8 = 71
    fmt.Println(v2, v28)
    v29 := GenS1{F0: 0}
    v3 = append(v3, 2147483647)
    fmt.Println(v2, v28)
    v30 := GenS1{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v9
}
