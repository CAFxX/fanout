package main

import (
    "fmt"
)


func main() {
    var v1 int = 42
    v2 := make(map[int]string)
    var v3 int16 = 32767
    v4 := []uint32{900, 604, 682, 693}
    v5 := make(map[string]bool)
    v6 := []int{v1}
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint = 2
            _ = len(v6)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint = 42
    var v11 uint = 65535
    fmt.Println(v11, v11)
    if (9223372036854775807 == 822) {
        v2[(^(v1))] = (("hello" + (("" + "") + ("\n\t" + "\n\t"))) + ((((("" + "a") + "hello") + (("a" + "hello") + "\n\t")) + "hello") + ("" + ("hello" + "foo bar"))))
    }
    delete(v5, ("a" + "a"))
    var v12 uint8 = 127
    var v13 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
