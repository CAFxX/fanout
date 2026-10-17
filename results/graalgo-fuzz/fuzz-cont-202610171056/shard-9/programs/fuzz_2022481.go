package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 int8 }
func genFn0(p0 int64, p1 string, p2 string) bool {
    switch (p2 + p1) {
    case "\n\t":
        if (1000 == 1) {
            p2 = p2
        } else {
            v6 := GenS1{F0: 133, F1: -610, F2: p2}
            _ = v6
        }
    case "":
        var v7 int = 1
        _ = v7
    }
    var v8 uint32 = 4294967295
    p2 = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (0 > 799)
}

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[int]bool)
    var v3 int16 = 42
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    if (("hello" < "hello") || (!false)) {
        for fi9 := 0; fi9 < 18; fi9++ {
            if true {
                v1 = append(v1, "hello")
            } else {
                delete(v5, fi9)
            }
            v3 -= 42
            _ = fi9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi10, fi10)
            if (!true) {
                if true {
                    v3 = v3
                    for fi11 := 0; fi11 < 4; fi11++ {
                        var v12 int = 128
                        _ = fi11
                        _ = v12
                    }
                    if true {
                        v13 := GenS1{F0: 44, F1: 872, F2: "hello"}
                        v14 := GenS2{F0: 1}
                        if true {
                            var v15 uint16 = 65535
                            _ = v15
                        }
                        _ = v13
                        _ = v14
                    } else {
                        for fi16 := 0; fi16 < 11; fi16++ {
                            v17 := v5[fi10]
                            v17 = (fi10 | (fi16 % (((fi16 % ((fi10) | 1))) | 1)))
                            fmt.Println(fi16, v17)
                            _ = fi16
                            _ = v17
                        }
                        var v18 int = fi10
                        for fi19 := 0; fi19 < 1; fi19++ {
                            switch "hello" {
                            case "foo bar":
                                v18 += (fi19 | 1)
                            case "hello":
                                v3 += 0
                            }
                            v18 = (fi10 >> ((870) & 63))
                            _ = fi19
                        }
                        _ = v18
                    }
                }
                v3 = v3
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 string = ((((("a" + "") + "foo bar") + ((((("" + "hello") + "hello") + ("" + "")) + "hello") + ((("a" + "\n\t") + "") + ("\n\t" + "\n\t")))) + "a") + "")
                    if (true || false) {
                        var v22 uint8 = 128
                        _ = v22
                    }
                    _ = len(v5)
                    _ = fi20
                    _ = v21
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
        }
    }
    if (0 == (v3 & -377)) {
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 int8 = -1
            v3 += v3
            _ = fi23
            _ = v24
        }
    } else {
        fmt.Println(v3)
        if true {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            v25 := GenS2{F0: 42}
            for fi26 := 0; fi26 < 3; fi26++ {
                v3 += 0
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi26
            }
            _ = v25
        }
    }
    v27 := GenS2{F0: 100}
    for fi28 := 0; fi28 < 1; fi28++ {
        v3 += 1
        _ = fi28
    }
    _ = len(v5)
    var v29 int32 = -432
    _ = len(v1)
    var v30 uint8 = 134
    v3 = v3
    v29 += (v29 % (((-(v29))) | 1))
    for fi31 := 0; fi31 < 12; fi31++ {
        v29 = v29
        _ = fi31
    }
    v32 := GenS2{F0: 1}
    switch 32767 {
    case -1:
        v30 += 0
    case 128:
        v33 := GenS2{F0: 42}
        _ = v33
    }
    var v34 string = "a"
    fmt.Println(v29, v30)
    if false {
        var v35 int = 32768
        v3 = v3
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
