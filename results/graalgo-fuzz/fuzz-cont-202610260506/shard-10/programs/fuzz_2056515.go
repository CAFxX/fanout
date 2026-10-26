package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8, p1 uint) int16 {
    fmt.Println(p1, p0)
    v5 := GenS1{F0: "a"}
    _ = p0
    _ = p1
    _ = v5
    return 1000
}

func main() {
    v1 := []string{((("a" + ("foo bar" + ("foo bar" + "hello"))) + "hello") + "\n\t"), ("a" + "a")}
    var v2 string = ("\n\t" + "\n\t")
    var v3 int8 = 34
    var v4 string = v2
    var v6 int16 = 486
    v2 = ((v2 + "\n\t") + ((("a" + (v4 + v2)) + (v2 + "a")) + v4))
    var v7 float64 = float64((1.5 + 3.14159))
    v3 += v3
    var v8 int64 = 42
    v4 = ("foo bar" + v2)
    if (v3 < int8(v8)) {
        var v9 string = (v4 + "a")
        _ = v9
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            _ = fi10
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 int16 = v6
        for fi13 := 0; fi13 < 19; fi13++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v14 uint = 125
            v1 = append(v1, v4)
            _ = fi13
            _ = v14
        }
        v15 := GenS2{F0: (42 != 312)}
        _ = fi11
        _ = v12
        _ = v15
    }
    v2 = v4
    var v16 uint = 1
    v2 = v2
    var v17 uint16 = 805
    v16 -= (^(v16))
    fmt.Println(v2)
    if true {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (-2.25 == (3.14159 + 3.14159)) {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    switch 17 {
                    case -343:
                        fmt.Println(v4, v8, v17)
                    case 327:
                        for fi19 := 0; fi19 < 13; fi19++ {
                            var v20 uint16 = (v17 << ((v16) & 63))
                            v4 = ("" + v4)
                            fmt.Println(v7)
                            _ = fi19
                            _ = v20
                        }
                    }
                    _ = fi18
                }
                v1 = append(v1, v4)
                v1 = append(v1, (("a" + v2) + (("hello" + v4) + v4)))
            }
            _ = len(v1)
        }
        if len(v1) > 0 {
            v1[0] = v2
        }
        v2 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
