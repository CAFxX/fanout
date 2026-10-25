package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int16 = 1
    v2 := []string{("\n\t" + "foo bar")}
    v3 := []string{}
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(fi6, fi7, fi7)
            var v8 int8 = -1
            _ = fi7
            _ = v8
        }
        fmt.Println(v1)
        _ = fi6
    }
    fmt.Println(v1)
    if len(v2) > 0 {
        v2[0] = ("\n\t" + ("a" + "a"))
    }
    v9 := GenS1{F0: 42}
    var v10 uint64 = 4294967295
    v1 -= 1
    for fi11 := 0; fi11 < 20; fi11++ {
        v4[("" + "\n\t")] = true
        v1 = v1
        var v12 uint32 = 65535
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v10)
    fmt.Println(v1)
    v13 := GenS1{F0: -2147483648}
    _ = len(v2)
    v14 := GenS1{F0: 2147483647}
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int8 = 0
            v1 = 1
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS1{F0: -2147483648}
    var v19 int16 = (v1 ^ 32767)
    var v20 uint = 0
    v19 = v1
    if false {
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 int = 1
            _ = len(v4)
            _ = fi21
            _ = v22
        }
        if (!(false && ((v1 > (v19 | -32768)) && false))) {
            switch -2147483648 {
            case 32768:
                for fi23 := 0; fi23 < 20; fi23++ {
                    if true {
                        for fi24 := 0; fi24 < 19; fi24++ {
                            var v25 float64 = 0.0
                            _ = fi24
                            _ = v25
                        }
                        if false {
                            var v26 bool = (65535 >= 604)
                            _ = v26
                        } else {
                            fmt.Println(fi23)
                            fmt.Println(v10, fi23, fi23)
                        }
                        v27 := GenS1{F0: -205}
                        _ = v27
                    } else {
                        fmt.Println(v1)
                    }
                    var v28 bool = true
                    v29 := GenS1{F0: 42}
                    _ = fi23
                    _ = v28
                    _ = v29
                }
            case 320:
                v20 = (v20 ^ v20)
            case 128:
                v20 = v20
            }
            v3 = append(v3, ("" + "foo bar"))
        } else {
            var v30 bool = true
            fmt.Println(v19, v10, v19)
            _ = v30
        }
    } else {
        v31 := GenS1{F0: -873}
        var v32 int32 = -1
        v19 += 0
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
