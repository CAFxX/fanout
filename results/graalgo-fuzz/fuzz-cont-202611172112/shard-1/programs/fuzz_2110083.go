package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []string{(("a" + "") + "\n\t"), ("" + ("foo bar" + ("" + "a"))), "", ("a" + ("\n\t" + "hello"))}
    var v2 uint8 = 99
    var v3 bool = true
    var v4 int = -570
    var v5 uint32 = 42
    var v6 uint64 = 42
    var v7 uint = 2
    if (((79 != 0) || v3) && (v3 && v3)) {
        v7 = v7
        switch v4 {
        case 0:
            var v8 uint = (^(v7))
            _ = v8
        case 42:
            v6 -= 613
        default:
            v9 := GenS1{F0: 709}
            _ = v9
        }
        v3 = true
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(fi10, v7)
            fmt.Println(v5)
            _ = fi11
        }
        switch ("hello" + "a") {
        case "":
            var v12 int32 = -833
            _ = v12
        case "a":
            v2 = 32
        default:
            fmt.Println(v5)
        }
        var v13 uint32 = uint32(v7)
        _ = fi10
        _ = v13
    }
    v14 := GenS1{F0: 0}
    if true {
        if len(v1) > 0 {
            v1[0] = ""
        }
        var v15 uint8 = (v2 ^ 128)
        var v16 uint32 = 416
        _ = v15
        _ = v16
    }
    if v3 {
        var v17 uint64 = (v6 | 711)
        _ = v17
    } else {
        if v3 {
            if (2147483648 > 2) {
                v18 := GenS1{F0: 2}
                _ = v18
            }
            if v3 {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v19 := GenS1{F0: 65535}
                if (v3 && (!v3)) {
                    v5 -= (v5 / ((v5) | 1))
                    v20 := GenS1{F0: 2}
                    _ = v20
                } else {
                    v21 := GenS1{F0: 37}
                    switch v4 {
                    case 0:
                        v3 = v3
                    default:
                        var v22 int = v4
                        _ = v22
                    }
                    _ = v21
                }
                _ = v19
            } else {
                if v3 {
                    fmt.Println(v2)
                    v23 := GenS1{F0: 0}
                    _ = v23
                } else {
                    var v24 uint64 = (v6 | (v6 << ((v7) & 63)))
                    v1 = append(v1, "a")
                    var v25 uint64 = (^(v24))
                    _ = v24
                    _ = v25
                }
                v1 = append(v1, (("foo bar" + "\n\t") + "foo bar"))
            }
            var v26 bool = (-745 <= 115)
            _ = v26
        } else {
            if (!v3) {
                fmt.Println(v2)
            } else {
                _ = len(v1)
            }
            _ = len(v1)
        }
    }
    v7 -= (v7 * (v7 + (v7 + v7)))
    if (-2147483648 > v4) {
        fmt.Println(v5, v7, v4)
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 int32 = 1
            _ = fi27
            _ = v28
        }
        if len(v1) > 0 {
            v1[0] = ""
        }
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    if (1 != 253) {
        var v29 int64 = 42
        _ = len(v1)
        v30 := GenS2{F0: -2147483648}
        _ = v29
        _ = v30
    }
    var v31 int32 = -1
    var v32 bool = (float32(1e10) != float32(0.1))
    fmt.Println(v5, v32, v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v32 {
        switch "a" {
        case "":
            _ = len(v1)
        case "hello":
            for fi33 := 0; fi33 < 19; fi33++ {
                for fi34 := 0; fi34 < 10; fi34++ {
                    v1 = append(v1, (("" + "foo bar") + (("" + "") + "")))
                    _ = fi34
                }
                for fi35 := 0; fi35 < 7; fi35++ {
                    fmt.Println(fi33, v5, v7)
                    v7 -= 0
                    v6 = (v6 & v6)
                    _ = fi35
                }
                var v36 string = "foo bar"
                _ = fi33
                _ = v36
            }
        case "a":
            for fi37 := 0; fi37 < 3; fi37++ {
                v38 := GenS2{F0: v31}
                switch "hello" {
                case "foo bar":
                    if len(v1) > 0 {
                        v1[0] = (((("a" + ("\n\t" + "foo bar")) + "\n\t") + ("foo bar" + "foo bar")) + "hello")
                    }
                }
                var v39 uint = v7
                _ = fi37
                _ = v38
                _ = v39
            }
        default:
            _ = len(v1)
        }
    } else {
        var v40 uint16 = 255
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v32, v31, v5)
        _ = v40
    }
    v6 = (v6 << (((v7 << (((^(v7))) & 63))) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
