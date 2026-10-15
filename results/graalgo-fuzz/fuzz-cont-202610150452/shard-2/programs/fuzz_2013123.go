package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := []string{"a", ("a" + "hello"), "", "hello"}
    var v5 uint64 = 706
    var v6 string = "hello"
    if (!(0.1 == 1.5)) {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(v1)
            var v8 string = (v6 + "\n\t")
            var v9 int64 = -1
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v2[2147483647] = (v6 + (("a" + v6) + "a"))
    for fi10 := 0; fi10 < 2; fi10++ {
        v6 = ((v6 + "hello") + "foo bar")
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint64 = 586
            _ = len(v4)
            v12 += v12
            _ = fi11
            _ = v12
        }
        v13 := v2[65535]
        _ = fi10
        _ = v13
    }
    fmt.Println(v1, v1)
    var v14 float32 = float32(1.5)
    for fi15 := 0; fi15 < 9; fi15++ {
        if (true && (423 >= 4294967295)) {
            v6 = ""
        } else {
            var v16 float32 = (3.14159 - 1.5)
            var v17 bool = false
            v16 += (1e10 - v14)
            _ = v16
            _ = v17
        }
        _ = fi15
    }
    v4 = append(v4, "a")
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int = (-(fi18))
        var v20 uint = uint(v5)
        _ = fi18
        _ = v19
        _ = v20
    }
    if (740 <= 554) {
        fmt.Println(v6)
        var v21 int = 32767
        _ = v21
    } else {
        _ = len(v4)
        v3 = (v6 == v6)
    }
    switch 1 {
    case 154:
        var v22 int64 = -274
        _ = v22
    case -311:
        v6 = "\n\t"
    }
    if len(v4) > 0 {
        v4[0] = ((((v6 + "\n\t") + v6) + v6) + v6)
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint8 = 128
        _ = fi23
        _ = v24
    }
    if false {
        var v25 uint64 = v5
        switch "foo bar" {
        case "foo bar":
            var v26 int16 = -10
            _ = v26
        case "":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            var v27 string = "a"
            _ = v27
        default:
            for fi28 := 0; fi28 < 9; fi28++ {
                var v29 string = v6
                v6 = v6
                _ = fi28
                _ = v29
            }
        }
        _ = v25
    } else {
        var v30 float64 = ((0.1 / (1.5 / (((-2.25 - 0.1) - -2.25) - float64(-2.25)))) - 1.5)
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
