package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 0
    v3 := make(map[string]string)
    var v4 uint64 = 1
    v5 := make(map[uint32]string)
    var v6 float64 = 3.14159
    switch -871 {
    case 128:
        var v7 float32 = 3.14159
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint16 = 65535
            var v10 uint32 = (^(v2))
            if (true && true) {
                fmt.Println(v6, v4, v4)
                for fi11 := 0; fi11 < 13; fi11++ {
                    v10 = (v10 / ((v10) | 1))
                    fmt.Println(v4, v2)
                    _ = fi11
                }
            } else {
                if true {
                    v2 += v2
                }
                v12 := []int{}
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v13 int16 = -1
    var v14 int32 = 255
    v2 += (^(v2))
    delete(v1, (("\n\t" + (("a" + "\n\t") + ("hello" + ("\n\t" + "hello")))) + ("" + "foo bar")))
    var v15 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
