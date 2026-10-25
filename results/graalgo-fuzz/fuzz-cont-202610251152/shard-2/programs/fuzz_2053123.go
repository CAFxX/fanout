package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]bool)
    var v3 int = 128
    v4 := make(map[int]int)
    fmt.Println(v1, v1)
    delete(v4, 1)
    fmt.Println(v3, v3)
    fmt.Println(v1, v1)
    var v5 int = v3
    var v6 string = "foo bar"
    v1 = 128
    v1 = 128
    v7 := GenS1{F0: -695}
    var v8 int32 = 1
    v9 := GenS1{F0: -949}
    var v10 int8 = -64
    var v11 uint8 = (v1 ^ 1)
    for fi12 := 0; fi12 < 11; fi12++ {
        if (true && (!(true || false))) {
            switch v3 {
            case -368:
                var v13 int16 = 1000
                _ = v13
            case 4294967295:
                switch v5 {
                case 699:
                    var v14 string = (("" + (v6 + ((v6 + ("\n\t" + v6)) + "hello"))) + "")
                    _ = v14
                case 42:
                    for fi15 := 0; fi15 < 13; fi15++ {
                        fmt.Println(v1)
                        v16 := GenS1{F0: -972}
                        var v17 float32 = 1e10
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
            }
            for fi18 := 0; fi18 < 10; fi18++ {
                if (true || true) {
                    var v19 uint = 946
                    _ = v19
                } else {
                    v20 := []int{403}
                    _ = v20
                }
                switch -2147483648 {
                case -1:
                    switch v5 {
                    case 0:
                        v8 -= v8
                    default:
                        var v21 uint = 2
                        _ = v21
                    }
                }
                _ = fi18
            }
            var v22 int64 = -1
            _ = v22
        }
        _ = fi12
    }
    if false {
        v1 += 127
        v23 := GenS1{F0: 42}
        delete(v2, 754)
        _ = v23
    } else {
        v1 += v11
    }
    switch (v3 * -1) {
    case 2147483647:
        var v24 uint8 = 42
        _ = v24
    case -302:
        v25 := []int{v3, 32767, v5, (v5 >> ((694) & 63))}
        _ = v25
    case 1:
        fmt.Println(v3, v10, v8)
    default:
        switch v6 {
        case "hello":
            v1 -= 128
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
