package main

import (
    "fmt"
)


func main() {
    v1 := []string{("a" + ("hello" + ("a" + "hello"))), ("foo bar" + "a")}
    v2 := []uint32{42, 2, 0}
    var v3 int32 = -781
    v3 = v3
    fmt.Println(v3)
    _ = len(v1)
    v3 = v3
    v3 = (^(v3))
    v3 -= (v3 / ((v3) | 1))
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(fi4, fi4)
            _ = fi5
        }
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 int16 = 0
            _ = fi6
            _ = v7
        }
        v2 = append(v2, 65535)
        _ = fi4
    }
    var v8 int8 = -128
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "\n\t")
    }
    var v9 string = "hello"
    switch 61 {
    case -605:
        fmt.Println(v9, v9)
    }
    if false {
        var v10 int8 = 100
        if (false || ((true && (!(-2.25 < 1.5))) || false)) {
            var v11 bool = (42 < 0)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v11
        } else {
            fmt.Println(v8)
        }
        _ = v10
    }
    v3 += (v3 >> ((897) & 63))
    var v12 uint32 = 622
    if (42 != 42) {
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint32 = v12
                _ = fi14
                _ = v15
            }
            v12 += (v12 >> ((1) & 63))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi13
        }
        var v16 uint32 = 1
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
