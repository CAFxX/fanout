package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int8 {
    v6 := []int{2147483647, 865, 205}
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint64 = 18446744073709551615
        var v9 bool = false
        _ = len(v6)
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 = append(v6, 32767)
    _ = v6
    return -128
}

func main() {
    var v1 int = 128
    v2 := []string{"foo bar", ("foo bar" + "foo bar"), "a", ("a" + ("a" + "\n\t"))}
    var v3 string = (("hello" + "a") + "hello")
    var v4 int8 = 60
    var v5 int64 = -1
    v5 = v5
    for fi10 := 0; fi10 < 6; fi10++ {
        switch 32767 {
        case 42:
            fmt.Println(fi10, v4, v4)
        default:
            v5 -= -1
        }
        var v11 uint8 = 42
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 1}
    var v13 bool = true
    if v13 {
        if (v13 || v13) {
            v3 = "a"
        } else {
            var v14 int = v1
            v1 = v1
            _ = v14
        }
        if v13 {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            fmt.Println(v4, v3)
            var v15 uint16 = 255
            _ = v15
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 string = (v3 + v3)
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint64 = 18446744073709551615
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: genFn0()}
    for fi21 := 0; fi21 < 6; fi21++ {
        v13 = (303 > 472)
        var v22 int32 = -32
        for fi23 := 0; fi23 < 19; fi23++ {
            v4 -= v4
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    v1 = -599
    _ = len(v2)
    for fi24 := 0; fi24 < 3; fi24++ {
        if (v13 && (!false)) {
            var v25 uint64 = 42
            var v26 bool = (v3 < v3)
            v27 := GenS1{F0: (v4 >> ((42) & 63))}
            _ = v25
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 int64 = v5
                _ = fi28
                _ = v29
            }
        }
        switch v1 {
        case 32767:
            switch "\n\t" {
            case "":
                v3 = (v3 + ((v3 + ("" + v3)) + ""))
            }
        case 1:
            var v30 uint = 18446744073709551615
            _ = v30
        }
        v3 = "hello"
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
