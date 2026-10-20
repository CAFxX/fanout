package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (0.1 - (-2.25 + 1e10))
    var v3 float32 = 1.5
    var v4 uint32 = 849
    fmt.Println(v4, v4, v4)
    for fi5 := 0; fi5 < 14; fi5++ {
        if (!(126 >= 25)) {
            v3 += (float32((1e10 / float64(((-2.25 / 1.5) * 0.1)))) + v3)
            var v6 float64 = v2
            _ = v6
        }
        var v7 float32 = v3
        _ = fi5
        _ = v7
    }
    v8 := v1["\n\t"]
    if (!false) {
        v8 += v8
        var v9 string = ""
        v10 := GenS1{F0: 42}
        _ = v10
        _ = v9
    }
    if (true && (-168 >= 523)) {
        v11 := GenS1{F0: 2}
        _ = v11
    } else {
        if false {
            var v12 uint32 = v4
            v2 += (v2 / v2)
            _ = v12
        }
    }
    switch 42 {
    case -698:
        fmt.Println(v2, v8, v3)
    case 128:
        var v13 int16 = 42
        _ = v13
    case 1:
        v2 = (v2 + (v2 / v2))
    }
    fmt.Println(v4, v8, v4)
    var v14 bool = true
    v15 := []int{v8, v8}
    var v16 string = ("\n\t" + "foo bar")
    var v17 uint16 = 42
    v14 = (!v14)
    v16 = v16
    for fi18 := 0; fi18 < 19; fi18++ {
        v17 = v17
        _ = fi18
    }
    var v19 int64 = 452
    v2 = 3.14159
    switch (^(v8)) {
    case 326:
        v8 += v8
    case -360:
        v14 = true
    case -1:
        v20 := GenS1{F0: 2}
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 float64 = v2
        if (!false) {
            switch (fi21 * (-(v8))) {
            case 1:
                v14 = (!false)
            }
            v23 := GenS1{F0: 18446744073709551615}
            for fi24 := 0; fi24 < 19; fi24++ {
                if v14 {
                    var v25 uint8 = 1
                    v26 := GenS1{F0: 688}
                    _ = v25
                    _ = v26
                }
                if v14 {
                    v22 += 3.14159
                    v2 += (float64(v3) * float64((float32(v22) - (v3 / v3))))
                }
                _ = fi24
            }
            _ = v23
        } else {
            var v27 int64 = 905
            v16 = (v16 + v16)
            for fi28 := 0; fi28 < 7; fi28++ {
                v27 = v19
                var v29 int8 = 1
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(v17, v3, v17)
            _ = fi30
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
