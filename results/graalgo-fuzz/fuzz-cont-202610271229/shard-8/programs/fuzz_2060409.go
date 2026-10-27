package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint16 = 758
    v3 := make(map[int]int)
    fmt.Println(v2)
    v1 = (v1 * 1.5)
    var v4 string = (("a" + ("a" + "\n\t")) + "\n\t")
    var v5 int64 = -474
    var v6 bool = (!false)
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 680}
        if false {
            var v9 int8 = -1
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if v6 {
        var v10 uint32 = 394
        var v11 int16 = 1
        _ = v10
        _ = v11
    }
    var v12 uint = 2
    delete(v3, 284)
    var v13 int16 = 0
    v6 = (!(-2.25 > 1e10))
    if (("" == (((("\n\t" + (v4 + ((v4 + v4) + "hello"))) + (v4 + v4)) + (("hello" + "a") + "\n\t")) + ("\n\t" + v4))) || v6) {
        fmt.Println(v6)
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: v13}
            _ = fi14
            _ = v15
        }
    } else {
        var v16 string = ("hello" + (((v4 + "hello") + "\n\t") + v4))
        switch "a" {
        case "":
            for fi17 := 0; fi17 < 14; fi17++ {
                for fi18 := 0; fi18 < 1; fi18++ {
                    if ((1 > 42) && ("\n\t" >= v16)) {
                        var v19 int8 = 127
                        _ = v19
                    } else {
                        v20 := []int{-9223372036854775808}
                        if (!((((42 < 42) || (!v6)) || true) && v6)) {
                            v21 := GenS1{F0: v13}
                            _ = v21
                        }
                        _ = len(v3)
                        _ = v20
                    }
                    var v22 string = ("\n\t" + v16)
                    _ = fi18
                    _ = v22
                }
                _ = fi17
            }
        }
        var v23 int32 = 826
        _ = v16
        _ = v23
    }
    v24 := []int{}
    if v6 {
        switch 432 {
        case 42:
            if len(v24) > 0 {
                v24[0] = -507
            }
        }
    } else {
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := GenS1{F0: (v13 / ((int16(v2)) | 1))}
            v27 := GenS1{F0: 1}
            var v28 int64 = v5
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        var v29 float32 = ((1e10 + 0.1) * (1.5 - (1e10 * 0.1)))
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
