package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 bool, p1 uint16, p2 int16) uint16 {
    v4 := []int{}
    if false {
        v5 := GenS1{F0: 42}
        _ = v5
    }
    if len(v4) > 0 {
        v4[0] = -1
    }
    p2 -= p2
    var v6 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return p1
}
func genFn1(p0 int8) uint64 {
    var v7 bool = ((42 < 1) || true)
    v8 := []int{}
    if v7 {
        switch -240 {
        case 981:
            var v9 int16 = 1
            _ = v9
        case -822:
            if (genFn0((true || ("" <= (("" + "") + ("foo bar" + "a")))), 32767, -370) >= 65535) {
                if v7 {
                    if true {
                        fmt.Println(p0, v7)
                        v10 := GenS1{F0: -1}
                        _ = v10
                    } else {
                        v11 := GenS1{F0: 9223372036854775807}
                        var v12 uint32 = 1
                        _ = v11
                        _ = v12
                    }
                    v7 = (51 != 0)
                } else {
                    fmt.Println(p0)
                }
                var v13 int8 = (^(p0))
                _ = len(v8)
                _ = v13
            }
        }
    } else {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: -1}
            _ = fi14
            _ = v15
        }
    }
    p0 = 34
    _ = len(v8)
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []string{("foo bar" + (("foo bar" + (("foo bar" + "a") + ((("" + "foo bar") + ("\n\t" + ("" + "hello"))) + ("\n\t" + ("foo bar" + "hello"))))) + "a")), ("a" + "a"), "", (("" + ("a" + ("foo bar" + "\n\t"))) + "")}
    var v2 float64 = float64(1e10)
    var v3 int32 = 2147483647
    fmt.Println(v2)
    if false {
        for fi16 := 0; fi16 < 14; fi16++ {
            if false {
                for fi17 := 0; fi17 < 5; fi17++ {
                    var v18 float32 = 1e10
                    var v19 string = ""
                    v2 = -2.25
                    var v20 string = "a"
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v21 float32 = 3.14159
                _ = v21
            } else {
                var v22 uint8 = 2
                fmt.Println(v22, v2)
                _ = v22
            }
            fmt.Println(v3)
            _ = fi16
        }
    }
    fmt.Println(v3, v3)
    var v23 int16 = 42
    v3 = -1
    fmt.Println(v23, v23)
    v24 := GenS1{F0: -533}
    fmt.Println(v23)
    if (v3 == v3) {
        var v25 uint16 = 301
        _ = v25
    }
    if true {
        _ = len(v1)
        var v26 float64 = -2.25
        var v27 int = 1
        _ = v26
        _ = v27
    }
    var v28 int = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
}
