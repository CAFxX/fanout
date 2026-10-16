package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 290
    var v3 bool = false
    switch -2147483648 {
    case 653:
        v3 = ((float64(0.1) * 3.14159) <= 1e10)
    }
    v4 := GenS1{F0: "foo bar", F1: (v2 / ((1) | 1)), F2: -32768}
    v3 = ((!((!v3) || v3)) && (!(false && (((225 <= 0) || v3) && true))))
    v5 := GenS1{F0: "\n\t", F1: (v2 ^ (v2 >> ((65535) & 63))), F2: 32767}
    switch ("hello" + "foo bar") {
    case "":
        v3 = v3
    case "hello":
        v3 = (1 > 0)
    default:
        v3 = ((v2 < 620) && v3)
    }
    v6 := []int{-525, 0}
    v2 -= (^(v2))
    fmt.Println(v2)
    switch 551 {
    case -9223372036854775808:
        v2 += v2
    case 289:
        v3 = ((!(!v3)) || true)
    default:
        v7 := GenS1{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 0, F2: -32768}
        _ = v7
    }
    v8 := GenS1{F0: ("hello" + "hello"), F1: v2, F2: -13}
    v2 += v2
    switch -9223372036854775808 {
    case 32767:
        _ = len(v6)
    case 4294967295:
        v2 = v2
    case -1:
        v3 = (!true)
    }
    var v9 uint64 = 18446744073709551615
    delete(v1, "")
    for fi10 := 0; fi10 < 18; fi10++ {
        if false {
            v11 := GenS1{F0: ("foo bar" + "foo bar"), F1: v2, F2: 0}
            v2 -= 0
            v3 = (v3 && (!((v3 || true) && v3)))
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            delete(v1, "")
            var v13 uint16 = 941
            _ = fi12
            _ = v13
        }
        _ = fi10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
